package entities

import (
	"github.com/carlosfgti/go-api/pkg/value_object"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       value_object.ID `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       value_object.NewId(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
