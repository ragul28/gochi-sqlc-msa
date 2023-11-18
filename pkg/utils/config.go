package utils

import "os"

// ConfigStore struct stores the flags
type ConfigStore struct {
	DB_URL string
	PORT   string
}

func ConfigEnv() ConfigStore {
	var cfg ConfigStore

	cfg.DB_URL = GetEnv("DB_URL", "postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	cfg.PORT = GetEnv("PORT", "8080")

	return cfg
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
