package db

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

func CreateConnection(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	slog.Info("Database connected to Postgres.", slog.String("package", "database"))

	return db, nil
}
