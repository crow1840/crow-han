package main

import (
	"crow-han/internal/app/user/handler"
	"crow-han/internal/conf"
	pb "crow-han/proto/user"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	_ "github.com/go-micro/plugins/v4/registry/consul"
	"github.com/kelseyhightower/envconfig"
	"go-micro.dev/v4/registry"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "user"
	version = "latest"
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

}

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
	if err := pb.RegisterUserHandler(srv.Server(), new(handler.User)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
