package user

import "github.com/gin-gonic/gin"

// RegisterRoutes registers all the user-related endpoints
func RegisterRoutes(r *gin.Engine) {
    r.POST("/users", CreateUser)
    r.GET("/users/:id", GetUser)
    r.PUT("/users/:id", UpdateUser)
    r.DELETE("/users/:id", DeleteUser)
}