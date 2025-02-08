package models

import (
	"gorm.io/gorm"
)

type Order struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Quantity  int
	OrderDate string  `gorm:"type:date;not null"`
	Product   Product `gorm:"foreignKey:ProductID"` // relasi ke model Product
}

func CreateOrder(db *gorm.DB, order *Order) error {
	return db.Create(order).Error
}

func GetOrderByID(db *gorm.DB, id uint) (*Order, error) {
	var order Order
	if err := db.Preload("Product").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func GetAllOrders(db *gorm.DB) ([]Order, error) {
	var orders []Order
	if err := db.Preload("Product").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
