package main

import (
	"log"
	"os"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"user-management/db"
	"user-management/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file on main: ", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db.Init(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to the database: ")
	}

	r := gin.Default()

	user.RegisterRoutes(r)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	log.Printf("Server running at %s:%s\n", host, port)
	err = r.Run(host + ":" + port)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
