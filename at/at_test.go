package at

import (
	"context"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	t.Run("returns an error if scheduled time is in the past", func(t *testing.T) {
		ts := time.Unix(0, 0)
		err := Schedule(context.TODO(), ts, func(ctx context.Context) {})
		if err != ErrTimeAlreadyPassed {
			t.Errorf("expected ErrTimeAlreadyPassed, got %+v", err)
		}
	})

	t.Run("successfully schedules a job", func(t *testing.T) {
		ts := new(time.Time)
		scheduledTime := time.Now().Add(2 * time.Second)

		fn := func(ts *time.Time) func(ctx context.Context) {
			return func(ctx context.Context) { *ts = time.Now() }
		}

		Schedule(context.TODO(), scheduledTime, fn(ts))
		time.Sleep(4 * time.Second)
		if ts.Second() != scheduledTime.Second() {
			t.Errorf("expected run time to be %d, got %d", scheduledTime.Unix(), ts.Unix())
		}
	})

	t.Run("successfully cancels a scheduled job before it executes", func(t *testing.T) {
		ts := new(time.Time)
		*ts = time.Now()
		scheduledTime := time.Now().Add(2 * time.Second)

		var b bool
		fn := func(ctx context.Context) {
			b = true
		}

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		cancel()
		Schedule(ctx, scheduledTime, fn)
		if b {
			t.Error("couldn't cancel scheduled job")
		}
	})
}
