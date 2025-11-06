package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/golang-gin/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine, db *[]models.User) {
	InitUserRouter(r, db)
	InitAuthRouter(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
