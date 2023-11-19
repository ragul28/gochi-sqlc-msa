package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
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
	dbconn.Ping()

	router := chi.NewRouter()
	// queries := sqlc.New(dbconn)

	slog.Info(fmt.Sprintf("Server started on Port: %s", cfg.PORT))
	err = http.ListenAndServe(":"+cfg.PORT, router)
	if err != nil {
		slog.Error("Error http server", err, slog.String("package", "main"))
	}
}
