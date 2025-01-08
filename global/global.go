package global

import (
	"database/sql"
	"github.com/sonha1/food_delivery_app/pkg/logger"
	"github.com/sonha1/food_delivery_app/pkg/setting"
)

var (
	Config = setting.Config{}
	Logger *logger.ZapLogger
	Mdbc   *sql.DB
)
