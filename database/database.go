package database

import (
	"JWT-Authentication/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DB_CONNECTION_STRING")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database ", err)
		return nil, err
	}
	fmt.Println("Connected to Database")
	return DB, nil
}

func MigrateDatabase() {
	DB, err := ConnectToDatabase()
	if err != nil {
		log.Fatal("There is error connecting to database ", err)
		return
	}
	DB.AutoMigrate(&models.User{})
}
