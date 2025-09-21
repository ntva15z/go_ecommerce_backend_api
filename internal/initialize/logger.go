package initialize

import (
	"github.com/ntva15z/go-ecommerce-backend-api/global"
	"github.com/ntva15z/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
