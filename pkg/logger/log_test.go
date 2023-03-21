package logger

import (
	"testing"
)

func init() {
	SetDebugLogger()
}

func TestInfo(t *testing.T) {
	Info("info message", 1)
	Debug("debug message", 2)
	Error("error message", 2)
}
