package setup

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"houkago-tea-time-api/pkg/api/v1"
)

var db *gorm.DB

func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to setup...")
	}

	// Create the db schema
	err = database.AutoMigrate(&v1.Product{})
	if err != nil {
		panic("Failed to migrate database...")
	}

	// Create the db
	database.Create(&v1.Product{
		ProductID: 100001,
		Name:      "Test Product",
		Cost:      9.99,
		Retail:    20.00,
	})

	db = database
	return db
}
