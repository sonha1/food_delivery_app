package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sonha1/food_delivery_app/internal/controllers"
)

type AuthRouter struct {
	c controllers.AuthController
}

func NewAuthRouter(c controllers.AuthController) *AuthRouter {
	return &AuthRouter{
		c: c,
	}
}

func (a *AuthRouter) SetupRouter(r *gin.RouterGroup) {
	r.Group("/auth")
	{
		r.POST("/login", a.c.Login)
		r.POST("/register")
	}

}
