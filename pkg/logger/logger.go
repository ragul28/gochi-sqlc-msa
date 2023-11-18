package logger

import (
	"log/slog"
	"os"
)

func InitLogger() {
	txtlogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(txtlogger)
}
