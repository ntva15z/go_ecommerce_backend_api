package initialize

import (
	"fmt"
	"time"

	"github.com/ntva15z/go-ecommerce-backend-api/global"
	"github.com/ntva15z/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	checkErrorPanic(err, "InitMysql Error")
	global.Logger.Info("InitMysql Successfully")
	global.Mdb = db

	// set pool
	SetPool()
	migration()
}

func SetPool() {
	m := global.Config.Mysql
	sblDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("mysql error $s: ", err)
	}
	sblDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
}

func migration() {
	err := global.Mdb.AutoMigrate(
		&po.Role{},
		&po.User{},
	)

	if err != nil {
		fmt.Println("Migration error: ", err)
	}
}
