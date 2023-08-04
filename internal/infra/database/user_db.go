package database

import (
	"github.com/carlosfgti/go-api/internal/entities"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entities.User) error {
	return u.DB.Create(user).Error
}
