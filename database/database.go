package database

import (
	"fmt"
	"github.com/frhnfrnk/go-books/config"
	"github.com/frhnfrnk/go-books/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() *gorm.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)

	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Author{})
	db.AutoMigrate(&models.Publisher{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})

	fmt.Println("Migrations completed.")
}
