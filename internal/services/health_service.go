package services

import "mix-online-api/internal/models"

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) HealthCheck() (models.GenericResponse, error) {
	return models.GenericResponse{
		Message: "Service is up and running 🚀",
		Status:  "SUCCESS",
		Code:    200,
	}, nil
}
