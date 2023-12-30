package userservice

import (
	"context"

	"github.com/ragul28/gochi-sqlc-msa/internal/user/dto"
	userrepository "github.com/ragul28/gochi-sqlc-msa/internal/user/repository"
)

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
}

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	return nil
}
