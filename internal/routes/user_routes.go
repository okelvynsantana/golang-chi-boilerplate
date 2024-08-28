package routes

import (
	"mix-online-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(userHandler *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/{userID}", userHandler.GetUser)

	return r
}
