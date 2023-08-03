package log

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/livebud/color"
)

type (
	Logger  = slog.Logger
	Handler = slog.Handler
	Level   = slog.Level
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

func Default() *Logger {
	return New(Filter(LevelInfo, &Console{
		Writer: os.Stderr,
		Color:  color.Default(),
	}))
}

// ParseLevel parses a string into a log level
func ParseLevel(level string) (Level, error) {
	switch level {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "warn":
		return LevelWarn, nil
	case "error":
		return LevelError, nil
	}
	return 0, fmt.Errorf("log: %q is not a valid level", level)
}

// New logger
func New(handler Handler) *Logger {
	return slog.New(handler)
}
