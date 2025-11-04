package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/golang-gin/controllers"
	"github.com/siddiq24/golang-gin/models"
	"github.com/siddiq24/golang-gin/repositories"
	"github.com/siddiq24/golang-gin/services"
)

func InitUserRouter(r *gin.Engine, users *[]models.User) {
	userRepo := repositories.NewUserRepository(users)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	usersRouter := r.Group("/users")
	{
		usersRouter.GET("", userController.GetAllUsers)
		usersRouter.GET("/:id", userController.GetUserById)
		usersRouter.POST("", userController.CreateUser)
		usersRouter.PATCH("/:id", userController.UpdateUser)
		usersRouter.DELETE("/:id", userController.DeleteUser)
	}
}