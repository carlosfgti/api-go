package database

import (
	"github.com/carlosfgti/go-api/internal/entities"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (prod *Product) Create(product *entities.Product) error {
	return prod.DB.Create(product).Error
}

func (prod *Product) FindByID(id string) (*entities.Product, error) {
	var product entities.Product
	err := prod.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (prod *Product) Update(product *entities.Product) error {
	_, err := prod.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return prod.DB.Save(product).Error
}

func (prod *Product) Delete(id string) error {
	product, err := prod.FindByID(id)
	if err != nil {
		return err
	}
	return prod.DB.Delete(product).Error
}

func (prod *Product) FindAll(page, limit int, sort string) ([]entities.Product, error) {
	var products []entities.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = prod.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = prod.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}
