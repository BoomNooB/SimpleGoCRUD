package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// local package
	"github.com/boomnoob/go-practice-sql/model"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("customer.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&model.Customers{})

	DB = db
}

func DisconnectDatabase() {
	if DB != nil {
		db, err := DB.DB()
		if err != nil {
			panic("Failed to get database connection")
		}

		err = db.Close() // This closes the underlying database connection
		if err != nil {
			panic("Failed to close database connection")
		}

		DB = nil
	}
}
