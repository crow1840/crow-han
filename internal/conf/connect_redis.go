package conf

import (
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
	"github.com/toolkits/pkg/logger"
)

var RedisCtx = context.Background()
var Rdb *redis.Client

func RedisConnect() {
	var myRedis redis.Options
	myRedis.Addr = MyEnvs.RidesIp + ":" + MyEnvs.RidesPort
	myRedis.Password = ""
	myRedis.DB = MyEnvs.RidesDbNum
	Rdb = redis.NewClient(&myRedis)
	logger.Infof("%+v", Rdb)
}
