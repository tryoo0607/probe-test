package server

import (
	"probe-test/config"
	"probe-test/util"
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
	util.WriteProbe("alive", false)

	ready.Store(false)
	util.WriteProbe("ready", false)

	started.Store(false)
	util.WriteProbe("startup", false)

	cfg := config.GetInstance()

	// liveness
	if cfg.ProbeDelayLiveness <= 0 {
		alive.Store(true)
		util.WriteProbe("alive", true)
	} else {
		time.AfterFunc(cfg.ProbeDelayLiveness, func() {
			alive.Store(true)
			util.WriteProbe("alive", true)
		})
	}

	// readiness
	if cfg.ProbeDelayReadiness <= 0 {
		ready.Store(true)
		util.WriteProbe("ready", true)
	} else {
		time.AfterFunc(cfg.ProbeDelayReadiness, func() {
			ready.Store(true)
			util.WriteProbe("ready", true)
		})
	}

	// startup
	if cfg.ProbeDelayStartup <= 0 {
		started.Store(true)
		util.WriteProbe("startup", true)
	} else {
		time.AfterFunc(cfg.ProbeDelayStartup, func() {
			started.Store(true)
			util.WriteProbe("startup", true)
		})
	}
}
