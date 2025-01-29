package db

import (
	"fmt"

	"log"
	"user-management/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(host, port, user, password, dbName string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	Migrate()
}

func Migrate() {
	if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
        log.Fatal("failed to create uuid-ossp extension: ", err)
    }
	
    err := DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("failed to migrate database: ", err)
    }
}
