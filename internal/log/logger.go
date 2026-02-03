package log

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func InitLogger(debug bool) {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}

	handler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      level,
		TimeFormat: "15:04:05",
	})

	slog.SetDefault(slog.New(handler))
}
