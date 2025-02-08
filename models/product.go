package models

import (
	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	Price       float64
	Category    string `gorm:"type:varchar(255);not null"`
}

func CreateProduct(db *gorm.DB, product *Product) error {
	return db.Create(product).Error
}

func GetProductByID(db *gorm.DB, id uint) (*Product, error) {
	var product Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func GetAllProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProduct(db *gorm.DB, id uint, updatedProduct *Product) error {
	return db.Model(&Product{}).Where("id = ?", id).Updates(updatedProduct).Error
}

func DeleteProduct(db *gorm.DB, id uint) error {
	return db.Delete(&Product{}, id).Error
}

func GetProductsByCategory(db *gorm.DB, category string) ([]Product, error) {
	var products []Product
	// Mengambil produk berdasarkan kategori
	if err := db.Where("category = ?", category).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
