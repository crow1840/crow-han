package main

import (
	gwServic "crow-han/internal/app/gate-way/service"
	"crow-han/internal/cache"
	"crow-han/internal/conf"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"go-micro.dev/v4/logger"
)

func init() {
	//初始化变量
	err := envconfig.Process("crow", &conf.MyEnvs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", &conf.MyEnvs)

	//初始化日志
	conf.LoggerInit()
	logger.Info("=============== micro ================")

	//启动micro框架
	go gwServic.RunMicro()

	//初始化redis
	//redis初始化
	conf.RedisConnect()
	cache.CheckRides()
}

func main() {
	logger.Info("=============== router ================")
	gwServic.ServerRouter()
}
