package server

import (
	"context"
	"time"
)

var startedAt = time.Now()

func timeoutCtx(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}

func sleepMs(ms int) { time.Sleep(time.Duration(ms) * time.Millisecond) }
