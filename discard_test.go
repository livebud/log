package log_test

import (
	"testing"

	"github.com/livebud/log"
)

func TestDiscard(t *testing.T) {
	log := log.Discard()
	log.Debug("hello", "args", 10)
	log.Field("planet", "world").Field("args", 10).Info("hello")
	log.Field("planet", "world").Field("args", 10).Warn("hello")
	log.Field("planet", "world").Field("args", 10).Error("hello world")
}
