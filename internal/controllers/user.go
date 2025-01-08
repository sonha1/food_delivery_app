package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sonha1/food_delivery_app/internal/services"
	"github.com/sonha1/food_delivery_app/pkg/response"
	"net/http"
)

type UserController interface {
	GetUser(ctx *gin.Context)
}
type userController struct {
	s services.UserService
}

func NewUserController(s services.UserService) UserController {
	return &userController{s: s}
}

func (c *userController) GetUser(ctx *gin.Context) {
	response.SuccessResponse(ctx, http.StatusOK, c.s.GetUser())
}
