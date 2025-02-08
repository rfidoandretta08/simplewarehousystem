package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	Quantity  int
	Location  string  `gorm:"type:varchar(255);not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

func CreateInventory(db *gorm.DB, inventory *Inventory) error {
	return db.Create(inventory).Error
}

func GetInventoryByProductID(db *gorm.DB, productID uint) (*Inventory, error) {
	var inventory Inventory
	//if err := db.Where("product_id = ?", productID).First(&inventory).Error; err != nil {

	//return nil, err
	//}
	if err := db.Preload("Product").Where("product_id = ?", productID).First(&inventory).Error; err != nil {
		return nil, err
	}
	return &inventory, nil

}

func UpdateInventory(db *gorm.DB, productID uint, updatedInventory *Inventory) error {
	return db.Model(&Inventory{}).Where("product_id = ?", productID).Updates(updatedInventory).Error
}
