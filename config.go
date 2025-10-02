package entp

import (
	"errors"
	"time"
)

const (
	defaultTimeout time.Duration = 2 * time.Second
)

// ErrAllServersUnavailable is returned when no NTP server responded successfully.
var ErrAllServersUnavailable = errors.New("all NTP servers are unavailable")

// Default ntp servers.
var defaultServers = []string{
	"time.cloudflare.com",
	"time.google.com",
	"pool.ntp.org",
	"time.windows.com",
}
