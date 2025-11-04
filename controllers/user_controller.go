package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/siddiq24/golang-gin/models"
	"github.com/siddiq24/golang-gin/services"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(serv *services.UserService) *UserController{
	return &UserController{Service: serv}
}

func (u *UserController) GetAllUsers(c *gin.Context){
	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		log.Println(err)
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: err.Error(),
		})
	}
	c.JSON(200, models.Ressponse{
		Success: true,
		Massage: "Berhasil mendapatkan semua users",
		Data: allUsers,
	})
}

func (u *UserController) GetUserById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "id tidak valid",
		})
		return
	}
	user, err := u.Service.GetUserById(id)
	if err != nil {
		log.Println(err)
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: err.Error(),
		})
		return
	}
	c.JSON(200, models.Ressponse{
		Success: true,
		Massage: "user ditemukan",
		Data: user,
	})
}

func (u *UserController) CreateUser(c *gin.Context){
	var newuser models.User
	err := c.ShouldBindJSON(&newuser)
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "bad request",
		})
		return
	}

	userNew, err := u.Service.CreateUser(newuser)
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: err.Error(),
		})
		return
	}
	c.JSON(200, models.Ressponse{
		Success: true,
		Massage: "berhasil menambahkan user",
		Data: userNew,
	})	
}

func (u *UserController) UpdateUser(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "param tidak valid",
		})
		return
	}

	newUser := models.User{Id: id}

	err = c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "bad request body",
		})
		return
	}

	err = u.Service.UpdateUser(&newUser)
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: err.Error(),
		})
		return
	}

	c.JSON(200, models.Ressponse{
		Success: true,
		Massage: "Berhasil mengupdate user",
		Data: newUser,
	})
}

func (u *UserController) DeleteUser(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "gagal mengambil param",
		})
		return
	}

	err = u.Service.DeleteUser(id)
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: err.Error(),
		})
		return
	}

	c.JSON(200, models.Ressponse{
		Success: true,
		Massage: fmt.Sprintf("user dengan id %d berhasil dihapus", id),
	})
}