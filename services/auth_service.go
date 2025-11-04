package services

import (
	"context"
	"errors"
	"github.com/siddiq24/golang-gin/models"
	"github.com/siddiq24/golang-gin/repositories"
)

type AuthService struct {
	userRepo *repositories.AuthRepository
}

func NewAuthService(userRepo *repositories.AuthRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}
func (s *AuthService) Register(req *models.User) (*models.User, error) {
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	user := &models.User{
		Nama:     req.Nama,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &models.User{
		Id:       user.Id,
		Nama:     user.Nama,
		Email:    user.Email,
		Password: req.Password,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *models.User) (*models.Ressponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &models.Ressponse{
		Success: true,
		Massage: "Berhasil Login",
		Data: models.User{
			Id:    user.Id,
			Nama:  user.Nama,
			Email: user.Email,
		},
	}, nil
}
