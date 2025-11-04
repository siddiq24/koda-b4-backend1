package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Ressponse struct {
	Success bool   `json:"success"`
	Massage string `json:"massage"`
	Data    any    `json:"data"`
}

type User struct {
	Id   int    `json:"id" validate:"gt"`
	Nama string `json:"nama" validate:"required,min=3,max=20"`
}

var Users []User = []User{
	{
		Id:   1,
		Nama: "Ayuningtyas",
	},
	{
		Id:   2,
		Nama: "Bunga Cahyaning",
	},
	{
		Id:   3,
		Nama: "Cintya Ayu Dewi",
	},
	{
		Id:   4,
		Nama: "Darmaruhi",
	},
	{
		Id:   5,
		Nama: "Elysia Berta",
	},
	{
		Id:   6,
		Nama: "Denanda Valencianita",
	},
}

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Ressponse{
			Success: true,
			Data:    Users,
		})
	})

	r.GET("users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, Ressponse{
				Success: false,
				Massage: "id tidak valid",
			})
			return
		}
		for i, user := range Users {
			if id == user.Id {
				ctx.JSON(200, Ressponse{
					Success: true,
					Data:    Users[i],
				})
				return
			}
			if i == (len(Users)-1){
				ctx.JSON(400, Ressponse{
					Success: false,
					Massage: "user not found!",
				})
				return
			}
		}

	})

	r.POST("/users", func(ctx *gin.Context) {
		var newuser User
		err := ctx.ShouldBindJSON(&newuser)
		if err != nil {
			ctx.JSON(400, Ressponse{
				Success: false,
				Massage: "bad request",
			})
			return
		}

		for i, user := range Users {
			if newuser.Id == user.Id || newuser.Nama == user.Nama {
				ctx.JSON(200, Ressponse{
					Success: false,
					Massage: "Id or user is exist",
					Data:    Users[i],
				})
				return
			}
		}
		Users = append(Users, newuser)


		ctx.JSON(200, Ressponse{
			Success: true,
			Massage: "Berhasil Menambahkan User",
			Data:    Users,
		})
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, Ressponse{
				Success: false,
				Massage: "param tidak valid",
			})
			return
		}

		var newUser User
		newUser.Id = id

		err = ctx.ShouldBind(&newUser)

		if err != nil {
			ctx.JSON(400, Ressponse{
				Success: false,
				Massage: "bad request",
			})
			return
		}

		for i, user := range Users {
			if user.Id == newUser.Id{
				Users[i] = newUser
				ctx.JSON(200, Ressponse{
					Success: true,
					Massage: "berhasil mengupdate user",
					Data: Users,
				})
				return
			}
			if i == (len(Users)-1){
				ctx.JSON(400, Ressponse{
					Success: false,
					Massage: "user not found!",
				})
				return
			}
		}
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, Ressponse{
				Success: false,
				Massage: "gagal mengambil param",
			})
			return
		}

		for i, user := range Users {
			if user.Id == id{
				Users = append(Users[:i], Users[i+1:]...)
				ctx.JSON(200, Ressponse{
					Success: true,
					Massage: "berhasil menghapus user",
					Data: Users,
				})
				return
			}
			if i == (len(Users)-1){
				ctx.JSON(400, Ressponse{
					Success: false,
					Massage: "user not found!",
				})
				return
			}
		}
	})

	r.Run(":8081")
}
