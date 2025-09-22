package initialize

import (
	"fmt"

	"github.com/ntva15z/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("mysql user name", global.Config.Mysql.Username)

	InitLogger()
	global.Logger.Info("config ok", zap.String("ok", "logger"))

	InitMysql()

	InitRedis()

	r := InitRouter()

	r.Run(":2222")
}
