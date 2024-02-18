package repository

import (
	"log"

	"shopping-cart-service/model"

	"gorm.io/gorm"
)

type CartRepo interface {
	Migrate() error
}

type cartRepo struct {
	DB *gorm.DB
}

func (c cartRepo) Migrate() error {
	log.Print("[CartRepository]...Migrate")
	return c.DB.AutoMigrate(&model.Cart{})
}

func NewCartRepo(db *gorm.DB) CartRepo {
	return cartRepo{
		DB: db,
	}
}
