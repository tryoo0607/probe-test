package util

import (
	"os"
	"path/filepath"

	"probe-test/config"
)

func WriteProbe(name string, v bool) {
	cfg := config.GetInstance()
	probeDir := cfg.ProbeDir

	_ = os.MkdirAll(probeDir, 0o755)
	val := "false"
	if v {
		val = "true"
	}
	_ = os.WriteFile(filepath.Join(probeDir, name), []byte(val), 0o644)
}
