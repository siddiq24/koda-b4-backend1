package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/golang-gin/controllers"
	"github.com/siddiq24/golang-gin/models"
	"github.com/siddiq24/golang-gin/repositories"
	"github.com/siddiq24/golang-gin/services"
)

func InitAuthRouter(r *gin.Engine, db *[]models.User) {
	userRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}
}
