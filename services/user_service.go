package services

import (
	"github.com/siddiq24/golang-gin/models"
	"github.com/siddiq24/golang-gin/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (u *UserService) GetAllUsers() (*[]models.User, error) {
	return u.Repo.GetAll()
}

func (u *UserService) GetUserById(id int) (models.User, error) {
	return u.Repo.GetById(id)
}

func (u *UserService) CreateUser(user models.User)(models.User, error){
	return u.Repo.Create(user)
}

func (u *UserService) UpdateUser(newUser *models.User)error{
	return  u.Repo.Update(newUser)
}

func (u *UserService) DeleteUser(id int) error {
	return u.Repo.Delete(id)
}