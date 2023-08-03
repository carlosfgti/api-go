package entities

import (
	"time"

	"github.com/carlosfgti/go-api/pkg/value_object"
)

type Product struct {
	ID        value_object.ID `json:"id"`
	Name      string          `json:"name"`
	Price     float64         `json:"price"`
	CreatedAt time.Time       `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        value_object.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	return product, nil
}
