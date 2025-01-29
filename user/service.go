package user

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "time"
    "user-management/models"
    "user-management/db"
)

// CreateUser handles the creation of a new user
func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.ID = uuid.New()
    user.CreatedAt = time.Now()

    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

// GetUser fetches a user by UUID
func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// UpdateUser updates the user's details
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Save(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// DeleteUser removes a user by UUID
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := db.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}