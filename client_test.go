package entp

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient_Now(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		client := NewClient()

		ntpTime, err := client.Now(ctx)

		assert.NoError(t, err)
		assert.NotZero(t, ntpTime)
	})

	t.Run("success_WithAddServers", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		client := NewClient(
			WithReplaceDefaultServers(),
			WithAddServers("0.europe.pool.ntp.org"),
		)

		ntpTime, err := client.Now(ctx)

		assert.NoError(t, err)
		assert.NotZero(t, ntpTime)
	})

	t.Run("error_ErrAllServersUnavailable", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		client := NewClient(
			WithReplaceDefaultServers("127.0.0.1"),
		)

		ntpTime, err := client.Now(ctx)

		assert.Equal(t, err, ErrAllServersUnavailable)
		assert.Equal(t, ntpTime, time.Time{})
	})

	t.Run("error_nonServers", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		client := NewClient(
			WithReplaceDefaultServers(),
		)

		ntpTime, err := client.Now(ctx)

		assert.Equal(t, err, ErrAllServersUnavailable)
		assert.Equal(t, ntpTime, time.Time{})
	})

	t.Run("error_timeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := NewClient(
			WithTimeout(1*time.Second),
			WithReplaceDefaultServers("127.0.0.1", "127.0.0.2", "127.0.0.3"),
		)

		ntpTime, err := client.Now(ctx)

		assert.Equal(t, err, ErrAllServersUnavailable)
		assert.Equal(t, ntpTime, time.Time{})
	})

	t.Run("error_contextCanceled", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		client := NewClient(
			WithTimeout(2*time.Second),
			WithReplaceDefaultServers("127.0.0.1", "127.0.0.2", "127.0.0.3"),
		)

		ntpTime, err := client.Now(ctx)

		assert.Equal(t, err, ctx.Err())
		assert.Equal(t, ntpTime, time.Time{})
	})
}
