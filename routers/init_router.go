package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/golang-gin/models"
)

func InitRouter(r *gin.Engine, db *[]models.User) {
	InitUserRouter(r, db)
	InitAuthRouter(r, db)
}