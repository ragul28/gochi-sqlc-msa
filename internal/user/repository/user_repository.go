package userrepository

import (
	"context"
	"database/sql"

	"github.com/ragul28/gochi-sqlc-msa/internal/db/sqlc"
	"github.com/ragul28/gochi-sqlc-msa/internal/user/entity"
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
	CreateUser(ctx context.Context, u *entity.UserEntity) error
	FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error)
	UpdateUser(ctx context.Context, u *entity.UserEntity) error
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) ([]entity.UserEntity, error)
	UpdatePassword(ctx context.Context, pass, id string) error
}

func (r *repository) CreateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) FindUserByID(ctx context.Context, id string) (*entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdateUser(ctx context.Context, u *entity.UserEntity) error {
	return nil
}

func (r *repository) DeleteUser(ctx context.Context, id string) error {
	return nil
}

func (r *repository) FindManyUsers(ctx context.Context) ([]entity.UserEntity, error) {
	return nil, nil
}

func (r *repository) UpdatePassword(ctx context.Context, pass, id string) error {
	return nil
}
