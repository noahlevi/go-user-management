package main

import (
    "log"
    "os"
    "path/filepath"
    
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "user-management/user"
    "user-management/db"
)

func main() {
    err := godotenv.Load(filepath.Join(".", ".env"))
    if err != nil {
        log.Fatal("Error loading .env file on main: ", err)
    }

    db.Init()

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