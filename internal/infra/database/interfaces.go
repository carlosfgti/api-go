package database

import "github.com/carlosfgti/go-api/internal/entities"

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}
