package models

type User struct {
	Id   int    `json:"id"`
	Nama string `json:"nama" binding:"required,min=3,max=20"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}