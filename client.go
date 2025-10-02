// Package entp provides NTP-based current time retrieval.
package entp

import (
	"context"
	"time"

	"github.com/beevik/ntp"
)

// Client holds NTP configuration.
type Client struct {
	Servers []string
	Timeout time.Duration
}

// Option configures Client.
type Option func(*Client)

// WithAddServers add "servers" to the end of the list of available default servers.
func WithAddServers(servers ...string) Option {
	return func(c *Client) {
		c.Servers = append(c.Servers, servers...)
	}
}

// WithReplaceDefaultServers replaces the list of default servers with "servers" completely.
func WithReplaceDefaultServers(servers ...string) Option {
	return func(c *Client) {
		c.Servers = servers
	}
}

// WithTimeout sets the per-request timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.Timeout = timeout
	}
}

// NewClient creates a Client with options.
func NewClient(opts ...Option) *Client {
	client := &Client{
		Servers: defaultServers,
		Timeout: defaultTimeout,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Now returns the current time using the first responding NTP server.
// It respects Client.Timeout and an optional ctx deadline (the tighter wins).
func (c *Client) Now(ctx context.Context) (time.Time, error) {
	if len(c.Servers) == 0 {
		return time.Time{}, ErrAllServersUnavailable
	}

	timeout := c.Timeout

	if deadline, ok := ctx.Deadline(); ok {
		if d := time.Until(deadline); d > 0 && (timeout <= 0 || d < timeout) {
			timeout = d
		}
	}

	opts := ntp.QueryOptions{
		Timeout: timeout,
	}

	for _, s := range c.Servers {
		select {
		case <-ctx.Done():
			return time.Time{}, ctx.Err()
		default:
		}

		resp, err := ntp.QueryWithOptions(s, opts)
		if err != nil {
			continue
		}

		if err := resp.Validate(); err != nil {
			continue
		}

		// ntp: current time = local now + server-estimated clock offset
		return time.Now().Add(resp.ClockOffset), nil
	}

	return time.Time{}, ErrAllServersUnavailable
}
