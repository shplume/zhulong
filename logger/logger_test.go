package logger

import "testing"

func TestNewLogger(t *testing.T) {
	logger := New()

	logger.Error("test for logger")
}
