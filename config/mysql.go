package config

import (
	"fmt"
	"mpc-api/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// database connection
func InitDB() {
	// declare struct config & variable connectionString

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	connectionString := os.Getenv("CONNECTION_DB") + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Println("Error initializing database:", err)
	}
}

// db migration
func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&Product{})
	/*
		TODO 2
		migrate struct product
	*/
}
