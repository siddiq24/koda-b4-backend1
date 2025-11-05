package main

import (
	"github.com/gin-gonic/gin"
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
	db := configs.InitDb()

	r := gin.Default()
	r.Use(middlewares.InitCorsMiddleware())
	routers.InitRouter(r, db)
	r.Run(":8081")
}
