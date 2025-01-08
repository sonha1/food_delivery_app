package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sonha1/food_delivery_app/internal/dto"
	"github.com/sonha1/food_delivery_app/internal/repo"
	"github.com/sonha1/food_delivery_app/pkg/response"
	"github.com/sonha1/food_delivery_app/utils/crypto"
)

type AuthService interface {
	Login(c *gin.Context, dto *dto.LoginRequest) (codeResults int, result *dto.LoginResponse, err error)
}

type authService struct {
	userRepo *repo.UserRepo
}

func NewAuthService(userRepo *repo.UserRepo) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Login(c *gin.Context, dto *dto.LoginRequest) (codeResults int, result *dto.LoginResponse, err error) {
	// find User
	user, err := s.userRepo.FindUserByUsername(dto.Username)
	if err != nil {
		return response.ErrCodeLoginFailed, result, fmt.Errorf("find user failed")
	} else if user == nil {
		return response.ErrCodeLoginFailed, result, fmt.Errorf("user not Existed")
	}

	// compare password
	isMatching, err := crypto.MatchingPassword(user.Password, dto.Password)
	if err != nil {
		return response.ErrCodeLoginFailed, result, fmt.Errorf("check password failed")
	} else if !isMatching {
		return response.ErrCodeLoginFailed, result, fmt.Errorf("password is incorrect")
	}

	return response.ErrCodeSuccess, nil, nil
}
