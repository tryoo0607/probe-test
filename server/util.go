package server

import (
	"context"
	"os"
	"runtime/debug"
	"time"
)

var startedAt = time.Now()

func readVersion() string {
	if v := os.Getenv("APP_VERSION"); v != "" {
		return v
	}
	if bi, ok := debug.ReadBuildInfo(); ok {
		return bi.Main.Version
	}
	return "dev"
}
func uptime() time.Duration { return time.Since(startedAt) }

func timeoutCtx(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}

func sleepMs(ms int) { time.Sleep(time.Duration(ms) * time.Millisecond) }
