package handler

import (
	"context"
	"crow-han/internal/app/auth/service"
	"go-micro.dev/v4/logger"

	pb "crow-han/proto/auth"
)

type Auth struct{}

func (e *Auth) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Auth.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Auth) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) (err error) {
	logger.Infof("Received Auth.Call request: %v", req)
	var loginRequest service.LoginInfo
	loginRequest.UserName = req.UserName
	loginRequest.Password = req.Password
	rsp.Token, err = service.Login(loginRequest)
	if err != nil {
		logger.Error(err)
		//rsp.ErrS = err.Error()
		return err
	}

	return nil
}

func (e *Auth) Verify(ctx context.Context, req *pb.VerifyRequest, rsp *pb.VerifyResponse) (err error) {
	logger.Infof("Received Auth.Call request: %v", req)
	var userId int
	rsp.Result, userId, rsp.UserName, rsp.UserRole, err = service.Verify(req.Token)
	rsp.UserId = int64(userId)
	return nil
}

func (e *Auth) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest, rsp *pb.RefreshTokenResponse) (err error) {
	logger.Infof("Received Auth.Call request: %v", req)
	rsp.NewToken, err = service.RefreshToken(req.Token)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
