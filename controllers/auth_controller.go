
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/golang-gin/models"
	"github.com/siddiq24/golang-gin/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	req := models.User{
		Nama: ctx.PostForm("nama"),
		Email: ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
	}

	newUser, err := c.authService.Register(&req)
	if err != nil {
		if err.Error() == "email already registered" {
			ctx.JSON(400, models.Ressponse{
				Success: false,
				Massage: err.Error(),
			})
			return
		}

		ctx.JSON(500, models.Ressponse{
			Success: false,
			Massage: "Gagal mendafatrkan user",
		})
		return
	}

	ctx.JSON(200, models.Ressponse{
		Success: true,
		Massage: "User registered successfully",
		Data: newUser,
	})
}
func (c *AuthController) Login(ctx *gin.Context) {
	req := models.User{
		Nama: ctx.PostForm("nama"),
		Email: ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
	}

	response, err := c.authService.Login(ctx.Request.Context(), &req)
	if err != nil {
		ctx.JSON(400, models.Ressponse{
			Success: false,
			Massage: err.Error(),
		})
		return
	}

	ctx.JSON(200, response)
}