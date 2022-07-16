package configs

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/topnarapat/go-crud/models"
)

var DB *gorm.DB

func Connection() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect database")
	}
	fmt.Println("Connect database success")

	db.AutoMigrate(&models.User{})

	DB = db
}
