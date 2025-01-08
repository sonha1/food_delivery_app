package initialize

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sonha1/food_delivery_app/global"
	"go.uber.org/zap"
	"time"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlc() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname))
	checkErrorPanic(err, "Init Mysql failed")
	global.Logger.Info("Init Mysql success")
	global.Mdbc = db

	SetPool()
}

func SetPool() {
	m := global.Config.Mysql
	global.Mdbc.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	global.Mdbc.SetMaxOpenConns(m.MaxOpenConns)
	global.Mdbc.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}
