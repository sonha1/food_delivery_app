package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/sonha1/food_delivery_app/global"
	"github.com/sonha1/food_delivery_app/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// rate limiter middleware

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	mainGroup := r.Group("/api/v1")
	{
		routers.RegisterRouter(mainGroup)
	}

	return r
}
