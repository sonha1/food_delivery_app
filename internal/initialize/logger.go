package initialize

import (
	"github.com/sonha1/food_delivery_app/global"
	"github.com/sonha1/food_delivery_app/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config)
}
