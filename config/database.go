package config

import (
	"log"

	"assigment/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/day23?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	// Melakukan AutoMigrate untuk membuat tabel berdasarkan model
	err = DB.AutoMigrate(&models.Product{}, &models.Inventory{}, &models.Order{})
	if err != nil {
		log.Fatal("Error during migration:", err)
	}
	log.Println("Tables migrated successfully")
}
