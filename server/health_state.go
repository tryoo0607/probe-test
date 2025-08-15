package server

import (
	"sync/atomic"
)

var (
	alive   atomic.Bool
	ready   atomic.Bool
	started atomic.Bool
)

func InitHealthState() {
	alive.Store(true)
	ready.Store(true)
	started.Store(true)
}
