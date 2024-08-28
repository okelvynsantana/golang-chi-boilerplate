package repositories

import (
	"database/sql"
	"errors"
	"mix-online-api/internal/models"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
