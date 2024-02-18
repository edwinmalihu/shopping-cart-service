package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);unique"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
}

func (Customer) TableName() string {
	return "customer"
}

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(255);unique;not null"`
	Description string  `json:"description" gorm:"type:varchar(255)"`
	Price       float64 `json:"price" gorm:"type:decimal(22,2)"`
	Stok        uint    `json:"stok"`
}

type Cart struct {
	gorm.Model
	CustomerID uint     `json:"customer_id"`
	ProductID  uint     `json:"product_id"`
	Quantity   uint     `json:"qty"`
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	Product    Product  `gorm:"foreignKey:ProductID"`
}

func (Product) TableName() string {
	return "product"
}

func (Cart) TableName() string {
	return "cart"
}
