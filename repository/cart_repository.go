package repository

import (
	"log"

	"shopping-cart-service/model"
	"shopping-cart-service/request"
	"shopping-cart-service/response"

	"gorm.io/gorm"
)

type CartRepo interface {
	Migrate() error
	AddCart(request.RequestAddCart) (model.Cart, error)
	ListCart() ([]response.ResponseCart, error)
	DeleteCart(uint) (model.Cart, error)
	DetailCart(uint) (response.ResponseCart, error)
	DetailProduct(uint) (model.Product, error)
	UpdateStok(uint, uint) (model.Product, error)
}

type cartRepo struct {
	DB *gorm.DB
}

// UpdateStok implements CartRepo.
func (c cartRepo) UpdateStok(id uint, stok uint) (data model.Product, err error) {
	return data, c.DB.Model(&data).Where("id = ?", id).Update("stok", stok).Error
}

// DetailProduct implements CartRepo.
func (c cartRepo) DetailProduct(req uint) (data model.Product, err error) {
	return data, c.DB.First(&data, "id = ? ", req).Error
}

// DetailCart implements CartRepo.
func (c cartRepo) DetailCart(req uint) (data response.ResponseCart, err error) {
	return data, c.DB.Raw(`select p.id as product_id, p."name" as name, p.price as price, c.quantity as quantity from cart as c join product as p on c.product_id = p.id where c.id = ? and c.deleted_at isnull`, req).Scan(&data).Error
}

// AddCart implements CartRepo.
func (c cartRepo) AddCart(req request.RequestAddCart) (data model.Cart, err error) {
	data = model.Cart{
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
		Quantity:   req.Quantity,
	}

	return data, c.DB.Create(&data).Error
}

// DeleteCart implements CartRepo.
func (c cartRepo) DeleteCart(req uint) (data model.Cart, err error) {
	return data, c.DB.Where("id = ? ", req).Delete(&data).Error
}

// ListCart implements CartRepo.
func (c cartRepo) ListCart() (data []response.ResponseCart, err error) {
	return data, c.DB.Raw(`select p.id, p."name", p.price , c.quantity from product as p join cart as c ON p.id = c.product_id where c.deleted_at isnull`).Scan(&data).Error
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
