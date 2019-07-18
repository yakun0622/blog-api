package log

import "testing"

func TestLogger(t *testing.T) {
	Logger.WithField("test-key", "test-value").Info("this is info")
	Logger.WithField("test-key", "test-value").Warn("this is warn")
	Logger.WithField("test-key", "test-value").Debug("this is debug")
	Logger.WithField("test-key", "test-value").Error("this is error")
}
