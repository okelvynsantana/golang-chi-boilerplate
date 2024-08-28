package server

import (
	"fmt"
	database "mix-online-api/internal/db"
	"mix-online-api/internal/routes"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	port int

	db database.Service
}

func NewServer(router *chi.Mux) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      SetupRouter(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server

}

func SetupRouter() *chi.Mux {
	db := database.New()
	router := chi.NewRouter()

	// Adiciona middlewares ao roteador
	router.Use(middleware.Logger)    // Middleware para logging das requisições
	router.Use(middleware.Recoverer) // Middleware para recuperação de panics

	// Configura as rotas de usuários
	routes.ConfigureRoutes(router, db.DBConn())

	return router
}
