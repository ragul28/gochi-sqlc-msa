package userservice

import userrepository "github.com/ragul28/gochi-sqlc-msa/internal/user/repository"

func NewUserService(repo userrepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userrepository.UserRepository
}

type UserService interface {
	CreateUser() error
}

func (s *service) CreateUser() error {
	return nil
}
