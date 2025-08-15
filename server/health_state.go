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
	changeState("alive", false)

	changeState("false", false)

	changeState("startup", false)

	cfg := config.GetInstance()

	// liveness
	if cfg.ProbeDelayLiveness <= 0 {
		changeState("alive", true)
	} else {
		time.AfterFunc(cfg.ProbeDelayLiveness, func() {
			changeState("alive", true)
		})
	}

	// readiness
	if cfg.ProbeDelayReadiness <= 0 {
		changeState("ready", true)
	} else {
		time.AfterFunc(cfg.ProbeDelayReadiness, func() {
			changeState("ready", true)
		})
	}

	// startup
	if cfg.ProbeDelayStartup <= 0 {
		changeState("startup", true)
	} else {
		time.AfterFunc(cfg.ProbeDelayStartup, func() {
			changeState("startup", true)
		})
	}
}

func changeState(name string, state bool) {
	started.Store(state)
	util.WriteProbe(name, state)
}
