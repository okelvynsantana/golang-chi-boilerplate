package handlers

import (
	"database/sql"
	"encoding/json"
	"mix-online-api/internal/repositories"
	"mix-online-api/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(db *sql.DB) *UserHandler {
	userRepo := &repositories.UserRepository{DB: db}
	userService := &services.UserService{Repo: userRepo}
	return &UserHandler{Service: userService}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		if err == services.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
