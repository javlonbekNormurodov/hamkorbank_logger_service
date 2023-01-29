package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	logger.Debug("Debug message", Any("debug", "ndfd"))
}

func TestInfo(t *testing.T) {
	logger.Info("Info message", Any("Info", "ndfd"))
}
