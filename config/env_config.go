package config

import (
	"os"
	"sync"
)

type Config struct {
	HTTPPort string
	TCPPort  string
	GRPCPort string
}

var (
	once sync.Once
	c    *Config
)

func Load() *Config {
	return &Config{
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		TCPPort:  getEnv("TCP_PORT", "50051"),
		GRPCPort: getEnv("GRPC_PORT", "50051"),
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
