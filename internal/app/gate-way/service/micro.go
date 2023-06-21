package service

import (
	"crow-han/internal/app/gate-way/handler"
	"crow-han/internal/conf"
	pb "crow-han/proto/gate-way"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "gate-way"
	version = "latest"
	srv     micro.Service
	//AuthSrv auth.AuthService
	//UserSrv user.UserService
)

//func init() {
//
//	//AuthSrv = auth.NewAuthService("auth", srv.Client())
//	//UserSrv = user.NewUserService("user", srv.Client())
//}

func RunMicro() {
	srv = micro.NewService()
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
	if err := pb.RegisterGateWayHandler(srv.Server(), new(handler.GateWay)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
