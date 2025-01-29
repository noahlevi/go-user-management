package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "user-management/models"
)

var DB *gorm.DB

func Init() {
    dsn := "host=localhost user=gouser dbname=yourdb password=yourpassword sslmode=disable"
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