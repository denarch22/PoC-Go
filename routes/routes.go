package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sistema-usuarios-go/controllers"
	"sistema-usuarios-go/middleware"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)

	auth := r.Group("/admin")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/users", controllers.CreateUser)
		auth.GET("/users", controllers.GetUsers)
	}

	return r
}
