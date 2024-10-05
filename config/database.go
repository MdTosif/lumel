package config

import (
	"log"

	"github.com/mdtosif/lumel/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func migrate(db *gorm.DB) {
    db.AutoMigrate(&models.Product{}, &models.Customer{}, &models.Payment{}, &models.Order{})
}

func InitDatabase() {
    dsn := "host=localhost user=root password=root dbname=lumel port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database", err)
    }
    log.Println("Database connection successful")

	migrate(DB)
}
