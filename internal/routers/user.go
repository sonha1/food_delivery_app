package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sonha1/food_delivery_app/internal/controllers"
)

type UserRouter struct {
	c controllers.UserController
}

func NewUserRouter(
	c controllers.UserController,
) *UserRouter {
	return &UserRouter{c: c}
}

func (r *UserRouter) SetupRouter(rg *gin.RouterGroup) {
	rg.Group("/user")
	{
		rg.GET("", r.c.GetUser)
	}
}
