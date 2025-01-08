package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sonha1/food_delivery_app/internal/dto"
	"github.com/sonha1/food_delivery_app/internal/services"
	"github.com/sonha1/food_delivery_app/pkg/response"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	s services.AuthService
}

func NewAuthController(s services.AuthService) AuthController {
	return &authController{
		s: s,
	}
}

func (ac *authController) Login(c *gin.Context) {
	body := &dto.LoginRequest{}
	if err := c.ShouldBindJSON(body); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
	}

	if err := body.Validate(); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
	}

	code, result, err := ac.s.Login(c, body)
	if err != nil {
		response.ErrorResponse(c, code, err.Error())
	}

	response.SuccessResponse(c, code, result)
}

func (ac *authController) Register(c *gin.Context) {

}
