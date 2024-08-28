package routes

import (
	"database/sql"
	"mix-online-api/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func ConfigureRoutes(router *chi.Mux, db *sql.DB) {
	userHandler := handlers.NewUserHandler(db)
	healthHandler := handlers.NewHealthHandler()
	router.Mount("/users", UserRoutes(userHandler))
	router.Mount("/health", HealthRoutes(healthHandler))
}
