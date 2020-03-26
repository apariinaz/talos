// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health/health.proto

package health

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HealthCheck_ServingStatus int32

const (
	HealthCheck_UNKNOWN     HealthCheck_ServingStatus = 0
	HealthCheck_SERVING     HealthCheck_ServingStatus = 1
	HealthCheck_NOT_SERVING HealthCheck_ServingStatus = 2
)

var HealthCheck_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}

var HealthCheck_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheck_ServingStatus) String() string {
	return proto.EnumName(HealthCheck_ServingStatus_name, int32(x))
}

func (HealthCheck_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e55f18a46a9fb2dc, []int{1, 0}
}

type ReadyCheck_ReadyStatus int32

const (
	ReadyCheck_UNKNOWN   ReadyCheck_ReadyStatus = 0
	ReadyCheck_READY     ReadyCheck_ReadyStatus = 1
	ReadyCheck_NOT_READY ReadyCheck_ReadyStatus = 2
)

var ReadyCheck_ReadyStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "READY",
	2: "NOT_READY",
}

var ReadyCheck_ReadyStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"READY":     1,
	"NOT_READY": 2,
}

func (x ReadyCheck_ReadyStatus) String() string {
	return proto.EnumName(ReadyCheck_ReadyStatus_name, int32(x))
}

func (ReadyCheck_ReadyStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e55f18a46a9fb2dc, []int{3, 0}
}

type HealthWatchRequest struct {
	IntervalSeconds      int64    `protobuf:"varint,1,opt,name=interval_seconds,json=intervalSeconds,proto3" json:"interval_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthWatchRequest) Reset()         { *m = HealthWatchRequest{} }
func (m *HealthWatchRequest) String() string { return proto.CompactTextString(m) }
func (*HealthWatchRequest) ProtoMessage()    {}
func (*HealthWatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55f18a46a9fb2dc, []int{0}
}

func (m *HealthWatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthWatchRequest.Unmarshal(m, b)
}

func (m *HealthWatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthWatchRequest.Marshal(b, m, deterministic)
}

func (m *HealthWatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthWatchRequest.Merge(m, src)
}

func (m *HealthWatchRequest) XXX_Size() int {
	return xxx_messageInfo_HealthWatchRequest.Size(m)
}

func (m *HealthWatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthWatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthWatchRequest proto.InternalMessageInfo

func (m *HealthWatchRequest) GetIntervalSeconds() int64 {
	if m != nil {
		return m.IntervalSeconds
	}
	return 0
}

type HealthCheck struct {
	Status               HealthCheck_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=health.HealthCheck_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55f18a46a9fb2dc, []int{1}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}

func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}

func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}

func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}

func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetStatus() HealthCheck_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheck_UNKNOWN
}

type HealthCheckResponse struct {
	Messages             []*HealthCheck `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55f18a46a9fb2dc, []int{2}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}

func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}

func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}

func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}

func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetMessages() []*HealthCheck {
	if m != nil {
		return m.Messages
	}
	return nil
}

type ReadyCheck struct {
	Status               ReadyCheck_ReadyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=health.ReadyCheck_ReadyStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ReadyCheck) Reset()         { *m = ReadyCheck{} }
func (m *ReadyCheck) String() string { return proto.CompactTextString(m) }
func (*ReadyCheck) ProtoMessage()    {}
func (*ReadyCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55f18a46a9fb2dc, []int{3}
}

func (m *ReadyCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadyCheck.Unmarshal(m, b)
}

func (m *ReadyCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadyCheck.Marshal(b, m, deterministic)
}

func (m *ReadyCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadyCheck.Merge(m, src)
}

func (m *ReadyCheck) XXX_Size() int {
	return xxx_messageInfo_ReadyCheck.Size(m)
}

func (m *ReadyCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadyCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ReadyCheck proto.InternalMessageInfo

func (m *ReadyCheck) GetStatus() ReadyCheck_ReadyStatus {
	if m != nil {
		return m.Status
	}
	return ReadyCheck_UNKNOWN
}

type ReadyCheckResponse struct {
	Messages             []*ReadyCheck `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReadyCheckResponse) Reset()         { *m = ReadyCheckResponse{} }
func (m *ReadyCheckResponse) String() string { return proto.CompactTextString(m) }
func (*ReadyCheckResponse) ProtoMessage()    {}
func (*ReadyCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e55f18a46a9fb2dc, []int{4}
}

func (m *ReadyCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadyCheckResponse.Unmarshal(m, b)
}

func (m *ReadyCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadyCheckResponse.Marshal(b, m, deterministic)
}

func (m *ReadyCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadyCheckResponse.Merge(m, src)
}

func (m *ReadyCheckResponse) XXX_Size() int {
	return xxx_messageInfo_ReadyCheckResponse.Size(m)
}

func (m *ReadyCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadyCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadyCheckResponse proto.InternalMessageInfo

func (m *ReadyCheckResponse) GetMessages() []*ReadyCheck {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterEnum("health.HealthCheck_ServingStatus", HealthCheck_ServingStatus_name, HealthCheck_ServingStatus_value)
	proto.RegisterEnum("health.ReadyCheck_ReadyStatus", ReadyCheck_ReadyStatus_name, ReadyCheck_ReadyStatus_value)
	proto.RegisterType((*HealthWatchRequest)(nil), "health.HealthWatchRequest")
	proto.RegisterType((*HealthCheck)(nil), "health.HealthCheck")
	proto.RegisterType((*HealthCheckResponse)(nil), "health.HealthCheckResponse")
	proto.RegisterType((*ReadyCheck)(nil), "health.ReadyCheck")
	proto.RegisterType((*ReadyCheckResponse)(nil), "health.ReadyCheckResponse")
}

func init() { proto.RegisterFile("health/health.proto", fileDescriptor_e55f18a46a9fb2dc) }

var fileDescriptor_e55f18a46a9fb2dc = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdb, 0xab, 0xd3, 0x40,
	0x10, 0xc6, 0xcd, 0x39, 0x24, 0x7a, 0x26, 0xd4, 0x86, 0x2d, 0x88, 0xa4, 0x20, 0x9a, 0xa7, 0x16,
	0x71, 0x23, 0x55, 0x04, 0x2f, 0x20, 0xd6, 0xd6, 0x0b, 0x42, 0x0a, 0x1b, 0xb5, 0xe8, 0x4b, 0xd9,
	0xa6, 0x6b, 0x12, 0x4c, 0xb2, 0xb1, 0xbb, 0x29, 0xd4, 0x67, 0xff, 0x37, 0xff, 0x2d, 0x49, 0x36,
	0x69, 0xd3, 0x8b, 0xe7, 0x29, 0x99, 0xc9, 0x37, 0x33, 0xdf, 0x2f, 0x7c, 0xd0, 0x8b, 0x18, 0x4d,
	0x64, 0xe4, 0xaa, 0x07, 0xce, 0xd7, 0x5c, 0x72, 0x64, 0xa8, 0xca, 0xee, 0x87, 0x9c, 0x87, 0x09,
	0x73, 0xab, 0xee, 0xb2, 0xf8, 0xe1, 0xb2, 0x34, 0x97, 0x5b, 0x25, 0x72, 0x5e, 0x03, 0xfa, 0x50,
	0xc9, 0xe6, 0x54, 0x06, 0x11, 0x61, 0xbf, 0x0a, 0x26, 0x24, 0x1a, 0x82, 0x15, 0x67, 0x92, 0xad,
	0x37, 0x34, 0x59, 0x08, 0x16, 0xf0, 0x6c, 0x25, 0xee, 0x6a, 0xf7, 0xb5, 0xc1, 0x25, 0xe9, 0x36,
	0x7d, 0x5f, 0xb5, 0x9d, 0x3f, 0x1a, 0x98, 0x6a, 0xc3, 0xdb, 0x88, 0x05, 0x3f, 0xd1, 0x73, 0x30,
	0x84, 0xa4, 0xb2, 0x50, 0x03, 0xb7, 0x47, 0x0f, 0x70, 0x6d, 0xaa, 0x25, 0xc2, 0x3e, 0x5b, 0x6f,
	0xe2, 0x2c, 0xf4, 0x2b, 0x21, 0xa9, 0x07, 0x9c, 0x17, 0xd0, 0x39, 0xf8, 0x80, 0x4c, 0xb8, 0xf9,
	0xc5, 0xfb, 0xe4, 0xcd, 0xe6, 0x9e, 0x75, 0xa3, 0x2c, 0xfc, 0x29, 0xf9, 0xfa, 0xd1, 0x7b, 0x6f,
	0x69, 0xa8, 0x0b, 0xa6, 0x37, 0xfb, 0xbc, 0x68, 0x1a, 0x17, 0xce, 0x3b, 0xe8, 0xb5, 0x0e, 0x10,
	0x26, 0x72, 0x9e, 0x09, 0x86, 0x5c, 0xb8, 0x95, 0x32, 0x21, 0x68, 0xc8, 0x4a, 0x3f, 0x97, 0x03,
	0x73, 0xd4, 0x3b, 0xe3, 0x87, 0xec, 0x44, 0xce, 0x6f, 0x00, 0xc2, 0xe8, 0x6a, 0xab, 0x60, 0x9e,
	0x1d, 0xc1, 0xdc, 0x6b, 0x86, 0xf7, 0x1a, 0xf5, 0x7a, 0x44, 0xf2, 0x14, 0xcc, 0x56, 0xfb, 0x90,
	0xe3, 0x0a, 0x74, 0x32, 0x7d, 0x33, 0xf9, 0x66, 0x69, 0xa8, 0x03, 0x57, 0x25, 0x85, 0x2a, 0x2f,
	0x9c, 0x09, 0xa0, 0xfd, 0xde, 0x1d, 0x02, 0x3e, 0x41, 0x40, 0xa7, 0x2e, 0xf6, 0x04, 0xa3, 0xbf,
	0x1a, 0x18, 0x8a, 0x0d, 0xbd, 0x02, 0x5d, 0x71, 0xdc, 0xc1, 0x2a, 0x03, 0xb8, 0xc9, 0x00, 0x9e,
	0x96, 0x19, 0xb0, 0xfb, 0xe7, 0x7e, 0x46, 0x73, 0x78, 0x0c, 0x7a, 0x15, 0x0a, 0x64, 0x1f, 0xaa,
	0xda, 0x49, 0xb9, 0x76, 0xc3, 0x63, 0x0d, 0xbd, 0x04, 0xbd, 0x32, 0xf9, 0x5f, 0x07, 0xf6, 0x19,
	0x96, 0x7a, 0x7c, 0xfc, 0xf0, 0xfb, 0x30, 0x8c, 0x65, 0x54, 0x2c, 0x71, 0xc0, 0x53, 0x57, 0xd2,
	0x84, 0x8b, 0x47, 0x62, 0x2b, 0x24, 0x4b, 0x85, 0xaa, 0x5c, 0x9a, 0xc7, 0x75, 0xe6, 0x97, 0x46,
	0xb5, 0xf8, 0xc9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xa0, 0x2a, 0x23, 0x0b, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Watch(ctx context.Context, in *HealthWatchRequest, opts ...grpc.CallOption) (Health_WatchClient, error)
	Ready(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadyCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/health.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) Watch(ctx context.Context, in *HealthWatchRequest, opts ...grpc.CallOption) (Health_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Health_serviceDesc.Streams[0], "/health.Health/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Health_WatchClient interface {
	Recv() (*HealthCheckResponse, error)
	grpc.ClientStream
}

type healthWatchClient struct {
	grpc.ClientStream
}

func (x *healthWatchClient) Recv() (*HealthCheckResponse, error) {
	m := new(HealthCheckResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthClient) Ready(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ReadyCheckResponse, error) {
	out := new(ReadyCheckResponse)
	err := c.cc.Invoke(ctx, "/health.Health/Ready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Check(context.Context, *empty.Empty) (*HealthCheckResponse, error)
	Watch(*HealthWatchRequest, Health_WatchServer) error
	Ready(context.Context, *empty.Empty) (*ReadyCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HealthWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthServer).Watch(m, &healthWatchServer{stream})
}

type Health_WatchServer interface {
	Send(*HealthCheckResponse) error
	grpc.ServerStream
}

type healthWatchServer struct {
	grpc.ServerStream
}

func (x *healthWatchServer) Send(m *HealthCheckResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Health_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Ready(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _Health_Ready_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Health_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "health/health.proto",
}