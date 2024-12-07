package setup

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"houkago-tea-time-api/pkg/api/v1"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to setup...")
	}

	err = database.AutoMigrate(&v1.Product{})
	if err != nil {
		return
	}

	DB = database
}
