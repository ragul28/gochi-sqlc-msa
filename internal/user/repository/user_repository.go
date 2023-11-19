package userrepository

import (
	"database/sql"

	"github.com/ragul28/gochi-sqlc-msa/internal/db/sqlc"
)

func NewUserRepository(db *sql.DB, q *sqlc.Queries) UserRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type UserRepository interface {
	CreateUser() error
}

func (r *repository) CreateUser() error {
	return nil
}
