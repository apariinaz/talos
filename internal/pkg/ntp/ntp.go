// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package ntp provides a time sync client via SNTP protocol.
package ntp

import (
	"bytes"
	"context"
	"fmt"
	"math/bits"
	"math/rand"
	"net"
	"reflect"
	"sync"
	"syscall"
	"time"

	"github.com/beevik/ntp"
	"github.com/u-root/u-root/pkg/rtc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/talos-systems/talos/internal/pkg/timex"
)

// Syncer performs time sync via NTP on schedule.
type Syncer struct {
	logger *zap.Logger

	timeServersMu  sync.Mutex
	timeServers    []string
	lastSyncServer string

	timeSyncNotified bool
	timeSynced       chan struct{}

	restartSyncCh chan struct{}
	epochChangeCh chan struct{}

	firstSync bool

	MinPoll, MaxPoll time.Duration

	// these functions are overridden in tests for mocking support
	CurrentTime CurrentTimeFunc
	NTPQuery    QueryFunc
	AdjustTime  AdjustTimeFunc
}

// NewSyncer creates new Syncer with default configuration.
func NewSyncer(logger *zap.Logger, timeServers []string) *Syncer {
	syncer := &Syncer{
		logger: logger,

		timeServers: append([]string(nil), timeServers...),
		timeSynced:  make(chan struct{}),

		restartSyncCh: make(chan struct{}, 1),
		epochChangeCh: make(chan struct{}, 1),

		firstSync: true,

		MinPoll: MinAllowablePoll,
		MaxPoll: MaxAllowablePoll,

		CurrentTime: time.Now,
		NTPQuery:    ntp.Query,
		AdjustTime:  timex.Adjtimex,
	}

	return syncer
}

// Synced returns a channel which is closed when time is in sync.
func (syncer *Syncer) Synced() <-chan struct{} {
	return syncer.timeSynced
}

// EpochChange returns a channel which receives a value each time jumps more than EpochLimit.
func (syncer *Syncer) EpochChange() <-chan struct{} {
	return syncer.epochChangeCh
}

func (syncer *Syncer) getTimeServers() []string {
	syncer.timeServersMu.Lock()
	defer syncer.timeServersMu.Unlock()

	return syncer.timeServers
}

func (syncer *Syncer) getLastSyncServer() string {
	syncer.timeServersMu.Lock()
	defer syncer.timeServersMu.Unlock()

	return syncer.lastSyncServer
}

func (syncer *Syncer) setLastSyncServer(lastSyncServer string) {
	syncer.timeServersMu.Lock()
	defer syncer.timeServersMu.Unlock()

	syncer.lastSyncServer = lastSyncServer
}

// SetTimeServers sets the list of time servers to use.
func (syncer *Syncer) SetTimeServers(timeServers []string) {
	syncer.timeServersMu.Lock()
	defer syncer.timeServersMu.Unlock()

	if reflect.DeepEqual(timeServers, syncer.timeServers) {
		return
	}

	syncer.timeServers = append([]string(nil), timeServers...)
	syncer.lastSyncServer = ""

	syncer.restartSync()
}

func (syncer *Syncer) restartSync() {
	select {
	case syncer.restartSyncCh <- struct{}{}:
	default:
	}
}

// Run runs the sync process.
//
// Run is usually run in a goroutine.
// When context is canceled, sync process aborts.
//
//nolint:gocyclo
func (syncer *Syncer) Run(ctx context.Context) {
	RTCClockInitialize.Do(func() {
		var err error

		RTCClock, err = rtc.OpenRTC()
		if err != nil {
			syncer.logger.Error("failure opening RTC, ignored", zap.Error(err))
		}
	})

	for {
		lastSyncServer, resp, err := syncer.query(ctx)
		if err != nil {
			return
		}

		// Set some variance with how frequently we poll ntp servers.
		// This is based on rand(MaxPoll) + MinPoll so we wait at least
		// MinPoll.
		nextPollInterval := time.Duration(rand.Intn(int(syncer.MaxPoll.Seconds())))*time.Second + syncer.MinPoll

		if resp == nil {
			// if no response was ever received, consider doing short sleep to retry sooner as it's not Kiss-o-Death response
			nextPollInterval = syncer.MinPoll / 2
		}

		if resp != nil && resp.Validate() == nil {
			err = syncer.adjustTime(resp.ClockOffset, resp.Leap, lastSyncServer, nextPollInterval)

			if err == nil {
				if !syncer.timeSyncNotified {
					// successful first time sync, notify about it
					close(syncer.timeSynced)

					syncer.timeSyncNotified = true
				}
			} else {
				syncer.logger.Error("error adjusting time", zap.Error(err))
			}
		}

		select {
		case <-ctx.Done():
			return
		case <-syncer.restartSyncCh:
			// time servers got changed, restart the loop immediately
		case <-time.After(nextPollInterval):
		}
	}
}

func (syncer *Syncer) query(ctx context.Context) (lastSyncServer string, resp *ntp.Response, err error) {
	lastSyncServer = syncer.getLastSyncServer()
	failedServer := ""

	if lastSyncServer != "" {
		resp, err = syncer.queryServer(lastSyncServer)
		if err != nil {
			syncer.logger.Error(fmt.Sprintf("ntp query error with server %q", lastSyncServer), zap.Error(err))

			failedServer = lastSyncServer
			lastSyncServer = ""
			err = nil
		}
	}

	if lastSyncServer == "" {
		var serverList []string

		serverList, err = syncer.resolveServers(ctx)
		if err != nil {
			return lastSyncServer, resp, err
		}

		for _, server := range serverList {
			if server == failedServer {
				// skip server which failed in previous sync to avoid sending requests with short interval
				continue
			}

			select {
			case <-ctx.Done():
				return lastSyncServer, resp, ctx.Err()
			case <-syncer.restartSyncCh:
				return lastSyncServer, resp, nil
			default:
			}

			resp, err = syncer.queryServer(server)
			if err != nil {
				syncer.logger.Error(fmt.Sprintf("ntp query error with server %q", server), zap.Error(err))
				err = nil
			} else {
				syncer.setLastSyncServer(server)
				lastSyncServer = server

				break
			}
		}
	}

	return lastSyncServer, resp, err
}

func (syncer *Syncer) resolveServers(ctx context.Context) ([]string, error) {
	var serverList []string

	for _, server := range syncer.getTimeServers() {
		ips, err := net.LookupIP(server)
		if err != nil {
			syncer.logger.Warn(fmt.Sprintf("failed looking up %q, ignored", server), zap.Error(err))
		}

		for _, ip := range ips {
			serverList = append(serverList, ip.String())
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	return serverList, nil
}

func (syncer *Syncer) queryServer(server string) (*ntp.Response, error) {
	resp, err := syncer.NTPQuery(server)
	if err != nil {
		return nil, err
	}

	if err = resp.Validate(); err != nil {
		return resp, err
	}

	return resp, err
}

// adjustTime adds an offset to the current time.
//
//nolint:gocyclo
func (syncer *Syncer) adjustTime(offset time.Duration, leapSecond ntp.LeapIndicator, server string, nextPollInterval time.Duration) error {
	var (
		buf  bytes.Buffer
		req  syscall.Timex
		jump bool
	)

	if offset < -AdjustTimeLimit || offset > AdjustTimeLimit {
		jump = true

		fmt.Fprintf(&buf, "adjusting time (jump) by %s via %s", offset, server)

		req = syscall.Timex{
			Modes: timex.ADJ_SETOFFSET | timex.ADJ_NANO | timex.ADJ_STATUS,
			Time: syscall.Timeval{
				Sec:  int64(offset / time.Second),
				Usec: int64(offset / time.Nanosecond % time.Second),
			},
		}

		// kernel wants tv_usec to be positive
		if req.Time.Usec < 0 {
			req.Time.Sec--
			req.Time.Usec += int64(time.Second / time.Nanosecond)
		}
	} else {
		fmt.Fprintf(&buf, "adjusting time (slew) by %s via %s", offset, server)

		pollSeconds := uint64(nextPollInterval / time.Second)
		log2iPollSeconds := 64 - bits.LeadingZeros64(pollSeconds)

		req = syscall.Timex{
			Modes:    timex.ADJ_OFFSET | timex.ADJ_NANO | timex.ADJ_STATUS | timex.ADJ_TIMECONST | timex.ADJ_MAXERROR | timex.ADJ_ESTERROR,
			Offset:   int64(offset / time.Nanosecond),
			Status:   timex.STA_PLL,
			Maxerror: 0,
			Esterror: 0,
			Constant: int64(log2iPollSeconds) - 4,
		}
	}

	switch leapSecond { //nolint:exhaustive
	case ntp.LeapAddSecond:
		req.Status |= timex.STA_INS
	case ntp.LeapDelSecond:
		req.Status |= timex.STA_DEL
	}

	logLevel := zapcore.DebugLevel

	if jump {
		logLevel = zapcore.InfoLevel
	}

	state, err := syncer.AdjustTime(&req)

	fmt.Fprintf(&buf, ", state %s, status %s", state, timex.Status(req.Status))

	if err != nil {
		logLevel = zapcore.WarnLevel

		fmt.Println(&buf, ", error was %s", err)
	}

	if syncer.firstSync && logLevel == zapcore.DebugLevel {
		// promote first sync to info level
		syncer.firstSync = false

		logLevel = zapcore.InfoLevel
	}

	if ce := syncer.logger.Check(logLevel, buf.String()); ce != nil {
		ce.Write()
	}

	if err == nil {
		if offset < -EpochLimit || offset > EpochLimit {
			// notify about epoch change
			select {
			case syncer.epochChangeCh <- struct{}{}:
			default:
			}
		}

		if jump {
			if RTCClock != nil {
				if rtcErr := RTCClock.Set(time.Now().Add(offset)); rtcErr != nil {
					syncer.logger.Error("error syncing RTC", zap.Error(rtcErr))
				} else {
					syncer.logger.Info("synchronized RTC with system clock")
				}
			}
		}
	}

	return err
}
