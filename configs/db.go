package configs

import "github.com/siddiq24/golang-gin/models"

func InitDb() *[]models.User{
	return &[]models.User{
		{
			Id:   1,
			Nama: "Ayuningtyas",
			Email: "ayu@koda.com",
			Password: "Ayu12345",
		},
		{
			Id:   2,
			Nama: "Bunga Cahyaning",
			Email: "bunga@koda.com",
			Password: "Bunga12345",
		},
		{
			Id:   3,
			Nama: "Cintya Ayu Dewi",
			Email: "cintya@koda.com",
			Password: "Cintya@koda.com",
		},
		{
			Id:   4,
			Nama: "Darmaruhi",
			Email: "darma@koda.com",
			Password: "Darma12345",
		},
		{
			Id:   5,
			Nama: "Elysia Berta",
			Email: "berta@koda.com",
			Password: "Berta12345",
		},
		{
			Id:   6,
			Nama: "Denanda Valencianita",
			Email: "denanda@koda.com",
			Password: "Denanda12345",
		},
	}
}