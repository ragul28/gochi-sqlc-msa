package main

import (
	"log/slog"

	"github.com/ragul28/gochi-sqlc-msa/internal/db"
	"github.com/ragul28/gochi-sqlc-msa/pkg/logger"
	"github.com/ragul28/gochi-sqlc-msa/pkg/utils"
)

func main() {
	logger.InitLogger()
	slog.Info("Starting user service")

	cfg := utils.ConfigEnv()

	dbconn, err := db.CreateConnection(cfg.DB_URL)
	if err != nil {
		slog.Error("Error connection to DB", "err", err, slog.String("package", "main"))
		return
	}
}
