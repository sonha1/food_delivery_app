package services

import "github.com/sonha1/food_delivery_app/internal/repo"

type UserService interface {
	GetUser() string
	GetAllUsers() []string
}

type userService struct {
	repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) UserService {
	return &userService{repo: repo}
}

func (u userService) GetUser() string {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetAllUsers() []string {
	//TODO implement me
	panic("implement me")
}
