package userservice

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/ragul28/gochi-sqlc-msa/internal/user/dto"
	"github.com/ragul28/gochi-sqlc-msa/internal/user/entity"
	userrepository "github.com/ragul28/gochi-sqlc-msa/internal/user/repository"
	"golang.org/x/crypto/bcrypt"
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
	UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error
	GetUserByID(ctx context.Context, id string) (*entity.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	FindManyUsers(ctx context.Context) (*entity.ManyUsersResponse, error)
	UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error
}

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	userExists, err := s.repo.FindUserByEmail(ctx, u.Email)
	if err != nil {
		slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
		return err
	}
	if userExists != nil {
		slog.Error("user already exists", slog.String("package", "userservice"))
		return errors.New("user already exists")
	}
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return errors.New("error to encrypt password")
	}
	newUser := entity.UserEntity{
		ID:        uuid.New().String(),
		Name:      u.Name,
		Email:     u.Email,
		Password:  string(passwordEncrypted),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = s.repo.CreateUser(ctx, &newUser)
	if err != nil {
		slog.Error("error to create user", "err", err, slog.String("package", "userservice"))
		return err
	}
	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	userExists, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}
	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user already exists")
	}
	var updateUser entity.UserEntity
	if u.Email != "" {
		verifyUserEmail, err := s.repo.FindUserByEmail(ctx, u.Email)
		if err != nil {
			slog.Error("error to search user by email", "err", err, slog.String("package", "userservice"))
			return err
		}
		if verifyUserEmail != nil {
			slog.Error("user already exists", slog.String("package", "userservice"))
			return errors.New("user already exists")
		}
		updateUser.Email = u.Email
	}
	updateUser.ID = id
	updateUser.Name = u.Name
	updateUser.UpdatedAt = time.Now()
	err = s.repo.UpdateUser(ctx, &updateUser)
	if err != nil {
		slog.Error("error to update user", "err", err, slog.String("package", "userservice"))
		return err
	}
	return nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*entity.UserResponse, error) {
	userExists, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return nil, err
	}
	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return nil, errors.New("user not found")
	}
	user := entity.UserResponse{
		ID:        userExists.ID,
		Name:      userExists.Name,
		Email:     userExists.Email,
		CreatedAt: userExists.CreatedAt,
		UpdatedAt: userExists.UpdatedAt,
	}
	return &user, nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
	userExists, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}
	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user not found")
	}
	err = s.repo.DeleteUser(ctx, id)
	if err != nil {
		slog.Error("error to delete user", "err", err, slog.String("package", "userservice"))
		return err
	}
	return nil
}

func (s *service) FindManyUsers(ctx context.Context) (*entity.ManyUsersResponse, error) {
	findManyUsers, err := s.repo.FindManyUsers(ctx)
	if err != nil {
		slog.Error("error to find many users", "err", err, slog.String("package", "userservice"))
		return nil, err
	}
	users := entity.ManyUsersResponse{}
	for _, user := range findManyUsers {
		userResponse := entity.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		users.Users = append(users.Users, userResponse)
	}
	return &users, nil
}

func (s *service) UpdateUserPassword(ctx context.Context, u *dto.UpdateUserPasswordDto, id string) error {
	userExists, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		slog.Error("error to search user by id", "err", err, slog.String("package", "userservice"))
		return err
	}
	if userExists == nil {
		slog.Error("user not found", slog.String("package", "userservice"))
		return errors.New("user not found")
	}
	// compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(u.OldPassword))
	if err != nil {
		slog.Error("invalid password", slog.String("package", "userservice"))
		return errors.New("invalid password")
	}
	// compare new password with password in database
	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(u.Password))
	if err == nil {
		slog.Error("new password is equal to old password", slog.String("package", "userservice"))
		return errors.New("new password is equal to old password")
	}
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		slog.Error("error to encrypt password", "err", err, slog.String("package", "userservice"))
		return errors.New("error to encrypt password")
	}
	err = s.repo.UpdatePassword(ctx, string(passwordEncrypted), id)
	if err != nil {
		slog.Error("error to update password", "err", err, slog.String("package", "userservice"))
		return err
	}
	return nil
}
