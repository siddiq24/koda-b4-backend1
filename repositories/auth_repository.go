package repositories

import (
	"errors"

	"github.com/siddiq24/golang-gin/models"
)

type AuthRepository struct {
	users  *[]models.User
}

func NewAuthRepository(user *[]models.User) *AuthRepository {
	return &AuthRepository{users: user}
}

func (a *AuthRepository) Create(user *models.User) error {
	nextId := 0
	for _, user := range *a.users{
		if user.Id > nextId{
			nextId = user.Id
		}
	}
	user.Id = nextId +1
	*a.users = append(*a.users, *user)
	return nil
}

func (a *AuthRepository) FindByEmail(email string) (*models.User, error) {
	for _, user := range *a.users {
		if user.Email == email {
			userCopy := user
			return &userCopy, nil
		}
	}

	return nil, errors.New("user not found")
}

func (a *AuthRepository) FindByID(id int) (*models.User, error) {
	for _, user := range *a.users {
		if user.Id == id {
			userCopy := user
			return &userCopy, nil
		}
	}

	return nil, errors.New("user not found")
}