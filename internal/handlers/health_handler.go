package handlers

import (
	"encoding/json"
	"mix-online-api/internal/services"
	"net/http"
)

type HealthHandler struct {
	service *services.HealthService
}

func NewHealthHandler() *HealthHandler {
	healthService := services.NewHealthService()
	return &HealthHandler{service: healthService}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response, err := h.service.HealthCheck()

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
