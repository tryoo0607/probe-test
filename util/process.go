package util

import (
	"context"
	"time"
)

var startedAt = time.Now()

func TimeoutCtx(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}

func SleepMs(ms int) { time.Sleep(time.Duration(ms) * time.Millisecond) }
