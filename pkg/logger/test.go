package logger

import "testing"

var (
	logger LoggerI
)

func TestMain(m *testing.M) {
	logger = NewLogger("loger_service", LevelInfo)
}
