package config

import (
	"os"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	HTTPPort string
	TCPPort  string
	GRPCPort string

	LivenessServiceName  string
	ReadinessServiceName string
	StartupServiceName   string

	ProbeDelay          time.Duration
	ProbeDelayLiveness  time.Duration
	ProbeDelayReadiness time.Duration
	ProbeDelayStartup   time.Duration
}

var (
	once sync.Once
	c    *Config
)

func Load() *Config {

	common := getEnvDurationSec("PROBE_DELAY_SEC", 0)

	return &Config{
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		TCPPort:  getEnv("TCP_PORT", "9090"),
		GRPCPort: getEnv("GRPC_PORT", "50051"),

		LivenessServiceName:  getEnv("LIVENESS_SERVICE_NAME", ""),
		ReadinessServiceName: getEnv("READINESS_SERVICE_NAME", "ready"),
		StartupServiceName:   getEnv("STARTUP_SERVICE_NAME", "startup"),

		ProbeDelay:          common,
		ProbeDelayLiveness:  getEnvDurationSecWithDefault("PROBE_DELAY_LIVENESS_SEC", common),
		ProbeDelayReadiness: getEnvDurationSecWithDefault("PROBE_DELAY_READINESS_SEC", common),
		ProbeDelayStartup:   getEnvDurationSecWithDefault("PROBE_DELAY_STARTUP_SEC", common),
	}
}

func GetInstance() *Config {
	once.Do(func() { c = Load() })
	return c
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

func getEnvDurationSec(key string, defSec int) time.Duration {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return time.Duration(n) * time.Second
		}
	}
	return time.Duration(defSec) * time.Second
}

func getEnvDurationSecWithDefault(key string, def time.Duration) time.Duration {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return time.Duration(n) * time.Second
		}
	}
	return def
}
