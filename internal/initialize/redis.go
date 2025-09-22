package initialize

import (
	"context"
	"fmt"

	"github.com/ntva15z/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	rc := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", rc.Host, rc.Port),
		Password: rc.Password, // no password set
		DB:       rc.Database, // use default DB
		PoolSize: 10,          // số lượng kết nối tối đa
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis init Error: ", zap.Error(err))
	}

	fmt.Println("Redis init is running")
	global.Rdb = rdb

	redisExample()

}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis setting", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis get setting", zap.Error(err))
		return
	}

	global.Logger.Info("value score is::", zap.String("score", value))
}
