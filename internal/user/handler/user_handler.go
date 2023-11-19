package userhandler

import (
	"net/http"

	userservice "github.com/ragul28/gochi-sqlc-msa/internal/user/service"
)

func NewUserHandler(service userservice.UserService) UserHandler {
	return &handler{
		service,
	}
}

type handler struct {
	service userservice.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {}
