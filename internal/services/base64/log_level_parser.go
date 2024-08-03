package services

import (
	"errors"
	"log/slog"
)

func ParseLogLevel(level string) (slog.Level, error) {
	levels := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	if val, ok := levels[level]; ok {
		return val, nil
	}
	return slog.LevelInfo, errors.New("incorrect Log Level")
}
