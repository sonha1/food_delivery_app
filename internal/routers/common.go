package routers

import "github.com/gin-gonic/gin"

type Router interface {
	SetupRouter(r *gin.RouterGroup)
}
