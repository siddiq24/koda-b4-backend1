package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/siddiq24/golang-gin/configs"
	"github.com/siddiq24/golang-gin/middlewares"
	"github.com/siddiq24/golang-gin/routers"
)

type Ressponse struct {
	Success bool   `json:"success"`
	Massage string `json:"massage"`
	Data    any    `json:"data"`
}

func main() {
	godotenv.Load()
	db := configs.InitDb()

	r := gin.Default()
	r.Use(middlewares.InitCorsMiddleware())
	routers.InitRouter(r, db)
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}
