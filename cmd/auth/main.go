package main

import (
	"crow-han/internal/app/auth/handler"
	"crow-han/internal/cache"
	"crow-han/internal/conf"
	pb "crow-han/proto/auth"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	_ "github.com/go-micro/plugins/v4/registry/consul"
	"github.com/kelseyhightower/envconfig"
	"go-micro.dev/v4/registry"

	"go-micro.dev/v4"
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
	//mysql初始化
	var tx conf.MyConnect
	tx.ConnectMysql()

	//初始化redis
	//redis初始化
	conf.RedisConnect()
	cache.CheckRides()
}

var (
	service = "auth"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()

	consulRegis := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			conf.MyEnvs.ConsulAddr,
		}
	})

	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consulRegis),
	)

	// Register handler
	if err := pb.RegisterAuthHandler(srv.Server(), new(handler.Auth)); err != nil {
		logger.Fatal(err)
	}

	////调用user
	//userSrv := user.NewUserService("user", srv.Client())
	//resp, err := userSrv.Call(context.Background(), &user.CallRequest{
	//	Name: "Liu Bei",
	//})
	//if err != nil {
	//	logger.Error(err)
	//}
	//logger.Infof("request : %+v", resp)
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
