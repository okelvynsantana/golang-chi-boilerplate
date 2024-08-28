package routes

import (
	"mix-online-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func HealthRoutes(healthHandler *handlers.HealthHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", healthHandler.HealthCheck)

	return r
}
