package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"user-management/models"
)

var DB *gorm.DB

func Init() {

	env_err := godotenv.Load("../.env")
	if env_err != nil {
		log.Fatal("Error loading .env file in connection")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	// Remove the AutoMigrate here
}

// Migrate function to run migrations
func Migrate() {
	err := DB.AutoMigrate(&models.User{}) // Use the models.User struct
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
}
