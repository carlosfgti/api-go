package database

import "gorm.io/gorm"

type Product struct {
	DB *gorm.DB
}
