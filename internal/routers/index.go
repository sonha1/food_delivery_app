package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sonha1/food_delivery_app/global"
	"github.com/sonha1/food_delivery_app/internal/controllers"
	"github.com/sonha1/food_delivery_app/internal/database"
	"github.com/sonha1/food_delivery_app/internal/repo"
	"github.com/sonha1/food_delivery_app/internal/services"
)

func RegisterRouter(r *gin.RouterGroup) {
	// query db
	q := database.New(global.Mdbc)

	// repo
	userRepo := repo.NewUserRepo(q)

	// service
	userSv := services.NewUserService(userRepo)
	authSv := services.NewAuthService(userRepo)
	// controller
	userController := controllers.NewUserController(userSv)
	authController := controllers.NewAuthController(authSv)

	// router
	userRouter := NewUserRouter(userController)
	authRouter := NewAuthRouter(authController)

	userRouter.SetupRouter(r)
	authRouter.SetupRouter(r)
}
