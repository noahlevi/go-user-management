package main

import (
    "log"
    "os"
    
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "user-management/user"
    "user-management/db"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database
    db.Init()

    // Create a new Gin router
    r := gin.Default()

    // Register user routes
    user.RegisterRoutes(r)

    // Get host and port from environment variables 
    host := os.Getenv("HOST")
    port := os.Getenv("PORT")

    // Run the server
    log.Printf("Server running at %s:%s\n", host, port)
    err = r.Run(host + ":" + port)
    if err != nil {
        log.Fatalf("Error starting the server: %v", err)
    }
}