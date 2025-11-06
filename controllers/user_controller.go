package controllers

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/siddiq24/golang-gin/models"
	"github.com/siddiq24/golang-gin/services"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(serv *services.UserService) *UserController {
	return &UserController{Service: serv}
}

// GetAllUsers godoc
// @Summary      Get all users with pagination and search
// @Description  Retrieve list of users with pagination support and optional search functionality
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page    query     int     false  "Page number"           default(1)
// @Param        search  query     string  false  "Search by name/email"  default("")
// @Success      200     {object}  models.Ressponse{data=object{users=[]models.User,page=int,total_pages=int,total_users=int,search=string}}  "Successfully retrieved users"
// @Failure      400     {object}  models.Ressponse  "Bad request"
// @Router       /users [get]
func (u *UserController) GetAllUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	search := c.DefaultQuery("search", "")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit := 5
	offset := (page - 1) * limit

	allUsers, total, err := u.Service.GetAllUsers(limit, offset, search)
	if err != nil {
		log.Println(err)
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: err.Error(),
		})
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(200, models.Ressponse{
		Success: true,
		Massage: "Get all users successfully",
		Data: gin.H{
			"users":       allUsers,
			"page":        page,
			"total_pages": totalPages,
			"total_users": total,
			"search":      search,
		},
	})
}

// GetUserById godoc
// @Summary      Get user by ID
// @Description  Get detailed information of a specific user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.Ressponse{data=models.User}  "User found successfully"
// @Failure      400  {object}  models.Ressponse  "Invalid user ID or user not found"
// @Router       /users/{id} [get]
func (u *UserController) GetUserById(c *gin.Context) {
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
		Data:    user,
	})
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the provided information via form data
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        name      formData  string  true   "User name"        minlength(3)
// @Param        email     formData  string  true   "User email"       format(email)
// @Param        password  formData  string  true   "User password"    minlength(6)
// @Param        age       formData  int     false  "User age"         minimum(0)  maximum(150)
// @Success      200       {object}  models.Ressponse{data=models.User}  "User created successfully"
// @Failure      400       {object}  models.Ressponse  "Bad request - validation error"
// @Router       /users [post]
func (u *UserController) CreateUser(c *gin.Context) {
	nama := c.PostForm("nama")
	email := c.PostForm("email")
	password := c.PostForm("password")
	ageStr := c.DefaultPostForm("age", "0")

	if nama == "" {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "nama wajib diisi",
		})
		return
	}

	if len(nama) < 3 {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "nama minimal 3 karakter",
		})
		return
	}

	if email == "" {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "email wajib diisi",
		})
		return
	}

	if password == "" {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "password wajib diisi",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "password minimal 6 karakter",
		})
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "age harus berupa angka",
		})
		return
	}

	if age < 0 || age > 150 {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "age tidak valid",
		})
		return
	}

	newuser := models.User{
		Nama:     nama,
		Email:    email,
		Password: password,
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
		Data:    userNew,
	})
}

// UpdateUser godoc
// @Summary      Update user information
// @Description  Update existing user information by ID via form data
// @Tags         users
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        id        	path      int     true   "User ID"
// @Param        nama      	formData  string  true   "User name"        minlength(3)
// @Param        email     	formData  string  true   "User email"       format(email)
// @Param        password  	formData  string  false  "User password"    minlength(6)
// @Param        image		formData  file    false  "User picture"
// @Success      200       	{object}  models.Ressponse{data=models.User}  "User updated successfully"
// @Failure      400       	{object}  models.Ressponse  "Bad request - invalid ID or validation error"
// @Router       /users/{id} [patch]
func (u *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "param id tidak valid",
		})
		return
	}

	nama := c.PostForm("nama")
	email := c.PostForm("email")
	password := c.PostForm("password")
	var filename string
	file, header, err := c.Request.FormFile("image")
	if err == nil {
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			c.JSON(400, models.Ressponse{Success: false, Massage: "file bukan gambar valid"})
			return
		}

		fileExt := filepath.Ext(header.Filename)
		filename = fmt.Sprintf("%s%s", nama, fileExt)
		os.MkdirAll("images/user", os.ModePerm)

		localPath := fmt.Sprintf("images/user/%s", filename)

		resized := imaging.Resize(img, 1000, 0, imaging.Lanczos)
		if err := imaging.Save(resized, localPath); err != nil {
			c.JSON(400, models.Ressponse{Success: false, Massage: "gagal menyimpan gambar"})
			return
		}
	}

	if nama == "" {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "nama wajib diisi",
		})
		return
	}

	if len(nama) < 3 {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "nama minimal 3 karakter",
		})
		return
	}

	if email == "" {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "email wajib diisi",
		})
		return
	}

	if password != "" && len(password) < 6 {
		c.JSON(400, models.Ressponse{
			Success: false,
			Massage: "password minimal 6 karakter",
		})
		return
	}

	newUser := models.User{
		Id:          id,
		Nama:        nama,
		Email:       email,
		Password:    password,
		ProfilePict: filename,
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
		Data:    newUser,
	})
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.Ressponse  "User deleted successfully"
// @Failure      400  {object}  models.Ressponse  "Bad request - invalid ID or user not found"
// @Router       /users/{id} [delete]
func (u *UserController) DeleteUser(c *gin.Context) {
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
