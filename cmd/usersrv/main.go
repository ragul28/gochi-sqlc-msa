package main

import (
	"log/slog"

	"github.com/ragul28/gochi-sqlc-msa/pkg/logger"
	"github.com/ragul28/gochi-sqlc-msa/pkg/utils"
)

func main() {
	logger.InitLogger()
	slog.Info("Starting user service")

	cfg := utils.ConfigEnv()

	slog.Info("Env config", slog.String("DB URL", cfg.DB_URL))
}
