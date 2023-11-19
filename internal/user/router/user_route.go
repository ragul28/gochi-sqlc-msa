package routes

import (
	"github.com/go-chi/chi"
	userhandler "github.com/ragul28/gochi-sqlc-msa/internal/user/handler"
)

func InitUserRoutes(router chi.Router, h userhandler.UserHandler) {
	router.Route("/user", func(r chi.Router) {
		r.Post("/", h.CreateUser)
	})
}
