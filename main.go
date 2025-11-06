package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/siddiq24/golang-gin/configs"
	_ "github.com/siddiq24/golang-gin/docs"
	"github.com/siddiq24/golang-gin/middlewares"
	"github.com/siddiq24/golang-gin/routers"
)

// @title           Backend User Management API
// @version         1.0
// @description     API for user management and authentication
// @host      		localhost:8085
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
// @schemes http https
// @tag.name authentication
// @tag.description Authentication endpoints (register, login, logout)
// @tag.name users
// @tag.description User management endpoints (CRUD operations)
// @BasePath  /
func main() {
	godotenv.Load()

	db := configs.InitDb()

	r := gin.Default()
	r.Use(middlewares.InitCorsMiddleware())
	routers.InitRouter(r, db)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}
