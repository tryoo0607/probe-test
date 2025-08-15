package server

import (
	"probe-test/config"
	"sync/atomic"
	"time"
)

var (
	alive   atomic.Bool
	ready   atomic.Bool
	started atomic.Bool
)

func InitHealthState() {
	alive.Store(false)
	ready.Store(false)
	started.Store(false)

	cfg := config.GetInstance()

	// liveness
	if cfg.ProbeDelayLiveness <= 0 {
		alive.Store(true)
	} else {
		time.AfterFunc(cfg.ProbeDelayLiveness, func() { alive.Store(true) })
	}

	// readiness
	if cfg.ProbeDelayReadiness <= 0 {
		ready.Store(true)
	} else {
		time.AfterFunc(cfg.ProbeDelayReadiness, func() { ready.Store(true) })
	}

	// startup
	if cfg.ProbeDelayStartup <= 0 {
		started.Store(true)
	} else {
		time.AfterFunc(cfg.ProbeDelayStartup, func() { started.Store(true) })
	}
}
