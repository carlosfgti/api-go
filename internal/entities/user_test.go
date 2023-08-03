package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Carlos", "carlos@especializati.com.br", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Carlos", user.Name)
	assert.Equal(t, "carlos@especializati.com.br", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Carlos", "carlos@especializati.com.br", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("0123456"))
}
