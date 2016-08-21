package at

import (
	"context"
	"errors"
	"time"
)

// ErrTimeAlreadyPassed is returned when a timestamp in the past is passed to
// Schedule
var ErrTimeAlreadyPassed = errors.New("at: can't schedule in the past")

// Schedule queues a job to run at the specific time
func Schedule(ctx context.Context, ts time.Time, fn func(ctx context.Context)) error {
	duration := ts.Sub(time.Now())
	if int64(duration) < 0 {
		return ErrTimeAlreadyPassed
	}

	go func(ctx context.Context, duration time.Duration) {
		select {
		case <-ctx.Done():
			// timed out / cancelled
		case <-time.After(duration):
			fn(ctx)
		}
	}(ctx, duration)
	return nil
}
