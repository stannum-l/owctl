package util

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func NewLogger(debug bool) *slog.Logger {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}

	opts := &tint.Options{
		Level:      level,
		TimeFormat: time.Kitchen,
	}

	logger := slog.New(tint.NewHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	return logger
}
