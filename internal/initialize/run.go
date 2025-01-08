package initialize

import "github.com/sonha1/food_delivery_app/global"

func Run() {
	LoadConfig()
	InitLogger()
	InitMysqlc()
	r := InitRouter()

	err := r.Run(":8080")
	if err != nil {
		global.Logger.Error("Start app error : " + err.Error())
		panic(err)
	}
}
