package userhandler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ragul28/gochi-sqlc-msa/internal/user/dto"
	userservice "github.com/ragul28/gochi-sqlc-msa/internal/user/service"
	httperr "github.com/ragul28/gochi-sqlc-msa/pkg/http_error"
	"github.com/ragul28/gochi-sqlc-msa/pkg/utils"
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

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "userhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := utils.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.service.CreateUser(r.Context(), req)
	if err != nil {
		slog.Error(fmt.Sprintf("error to create user: %v", err), slog.String("package", "userhandler"))
		if err.Error() == "cep not found" {
			w.WriteHeader(http.StatusNotFound)
			msg := httperr.NewNotFoundError("cep not found")
			json.NewEncoder(w).Encode(msg)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		msg := httperr.NewBadRequestError("error to create user")
		json.NewEncoder(w).Encode(msg)
		return
	}
}
