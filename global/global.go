package global

import (
	"github.com/ntva15z/go-ecommerce-backend-api/pkg/logger"
	"github.com/ntva15z/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
