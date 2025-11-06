package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/golang-gin/dto"
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

// Register godoc
// @Sumary 				Register for new user
// @Description 		Create a new user
// @Tags				Auth
// @Accept				x-www-form-urlencoded
// @Produce 			json
// @Param 				nama		formData	string	true "User Name"
// @Param				email		formData	string	true "User Email"
// @Param				password	formData	string	true "User Password"
// @Success 			200	{object}	models.Ressponse
// @Failur				400 {object}	models.Ressponse
// @Router				/auth/register	[post]
func (c *AuthController) Register(ctx *gin.Context) {
	req := dto.Register_Request{
		Nama:     ctx.PostForm("nama"),
		Email:    ctx.PostForm("email"),
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
	}

	ctx.JSON(200, models.Ressponse{
		Success: true,
		Massage: "User registered successfully",
		Data:    newUser,
	})
}

// Login godoc
// @Sumary				Login User
// @Description 		Authenticate user using email and password
// @Tags				Auth
// @Accept				x-www-form-urlencoded
// @Produce				json
// @Param				email		formData	string	true	"User Email"
// @Param				password	formData	string	true	"User Password"
// @Success				200 {object}	models.Ressponse
// @Failur				400	{object}	models.Ressponse
// @Router				/auth/login	[post]
func (c *AuthController) Login(ctx *gin.Context) {
	req := models.User{
		Nama:     ctx.PostForm("nama"),
		Email:    ctx.PostForm("email"),
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
