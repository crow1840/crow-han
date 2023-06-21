package handler

import (
	"context"
	"crow-han/internal/app/user/service"
	pb "crow-han/proto/user"
	"github.com/toolkits/pkg/logger"
)

func (e *User) UpdateUserSelfInfo(ctx context.Context, req *pb.UpdateUserSelfInfoRequest, rsp *pb.UpdateUserSelfInfoResponse) error {
	logger.Infof("Received User.GetUser request: %v", req)
	var updateUserSelfInfoRequire service.UpdateUserSelfInfoRequire
	updateUserSelfInfoRequire.UserID = req.UserId
	updateUserSelfInfoRequire.Email = req.Email
	updateUserSelfInfoRequire.Phone = req.Phone
	err := service.UpdateUserSelfInfo(updateUserSelfInfoRequire)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (e *User) UpdateUserSelfPassword(ctx context.Context, req *pb.UpdateUserSelfPasswordRequest, rsp *pb.UpdateUserSelfPasswordResponse) error {
	logger.Infof("Received User.GetUser request: %v", req)
	var updateUserPasswordRequire service.UpdateUserSelfPasswordRequire
	updateUserPasswordRequire.UserId = req.UserId
	updateUserPasswordRequire.UserName = req.UserName
	updateUserPasswordRequire.Password = req.Password
	updateUserPasswordRequire.NewPassword = req.NewPassword
	err := service.UpdateUserSelfPassword(updateUserPasswordRequire)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (e *User) GetUserSelfInfo(ctx context.Context, req *pb.GetUserSelfInfoRequest, rsp *pb.GetUserSelfInfoResponse) error {
	logger.Infof("Received User.GetUser request: %v", req)
	userInfo, err := service.GetUserSelfInfo(req.UserName)
	if err != nil {
		logger.Error(err)
		return err
	}
	rsp.ID = userInfo.ID
	rsp.UserName = userInfo.UserName
	rsp.Role = userInfo.Role
	rsp.Email = userInfo.Email
	rsp.Phone = userInfo.Phone
	return nil
}
