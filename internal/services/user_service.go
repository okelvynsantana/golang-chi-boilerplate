package services

import (
	"database/sql"
	"errors"
	"mix-online-api/internal/repositories"
)

var ErrUserNotFound = errors.New("user not found")

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{Repo: &repositories.UserRepository{DB: db}}
}

func (s *UserService) GetUserByID(id string) (interface{}, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		if err == repositories.ErrUserNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
