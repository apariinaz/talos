// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package v1alpha1

import (
	"os"
	"path/filepath"
	"strings"
)

// Namespace implements config.ImageCache interface.
func (c *ClusterImageCache) Namespace() string {
	return c.CacheNamespace
}

// Path implements the config.ImagesCache interface.
func (c *ClusterImageCache) Path() string {
	// Check if provided path is in /var or not.
	// If not, returned a paht relative to /var
	path := c.CachePath

	subs := strings.Split(
		strings.TrimLeft(c.CachePath, "/"),
		string(os.PathSeparator),
	)

	if subs[0] != "var" {
		path = filepath.Join("/var", c.CachePath)
	}

	return path
}
