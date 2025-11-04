package repositories

import (
	"errors"

	"github.com/siddiq24/golang-gin/models"
)

type UserRepository struct {
	Users *[]models.User
}

func NewUserRepository(users *[]models.User) *UserRepository {
	return &UserRepository{Users: users}
}

func (u *UserRepository) GetAll() (*[]models.User, error) {
	if len(*u.Users) == 0 {
		return nil, errors.New("belum ada user yang terdaftar")
	}
	return u.Users, nil
}

func (u *UserRepository) GetById(id int) (models.User, error) {
	for _, user := range *u.Users {
		if user.Id == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (u *UserRepository) Create(newUser models.User) (models.User, error) {
	var id int
	for _, user := range *u.Users {
		if user.Id > id {
			id = user.Id
		}
		if user.Email == newUser.Email{
			return models.User{}, errors.New("email telah terdaftar")
		}
	}
	newUser.Id = id + 1
	*u.Users = append(*u.Users, newUser)
	return newUser, nil
}

func (u *UserRepository) Update(newUser *models.User) error{
	for i, user := range *u.Users {
		if user.Id == newUser.Id {
			(*u.Users)[i] = *newUser
			return nil
		}
	}
	return errors.New("user not found")
}

func (u *UserRepository) Delete(id int) error{
	for i , user := range *u.Users {
		if user.Id == id {
			*u.Users = append((*u.Users)[:i], (*u.Users)[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}