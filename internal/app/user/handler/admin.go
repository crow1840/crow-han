package handler

import (
	"context"
	"crow-han/internal/app/user/service"
	"crow-han/proto/user"
	pb "crow-han/proto/user"
	"go-micro.dev/v4/logger"
)

type User struct{}

func (e *User) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received User.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *User) CreateUser(ctx context.Context, req *pb.CreateUserRequest, rsp *pb.CreateUserResponse) (err error) {
	logger.Infof("Received User.Call request: %v", req)
	var createUserRequire service.CreateUserRequire
	createUserRequire.UserName = req.UserName
	createUserRequire.Password = req.Password
	createUserRequire.Email = req.Email
	createUserRequire.Phone = req.Phone
	createUserRequire.Role = req.Role
	rsp.UserId, err = service.CreateUser(createUserRequire)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (e *User) ResetUserPassword(ctx context.Context, req *pb.ResetUserUserRequest, rsp *pb.ResetUserResponse) error {
	logger.Infof("Received User.Call request: %v", req)
	var resetUserPasswordRequire service.ResetUserPasswordRequire
	resetUserPasswordRequire.UserId = req.UserId
	resetUserPasswordRequire.NewPassword = req.NewPassword
	err := service.ResetUserPassword(resetUserPasswordRequire)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (e *User) GetUsers(ctx context.Context, req *pb.GetUsersRequest, rsp *pb.GetUsersResponse) error {
	logger.Infof("Received User.Call req: %v", req)
	var getUsersRequest service.GetUsersRequire
	getUsersRequest.PageNum = req.PageNum
	getUsersRequest.PageSize = req.PageSize
	getUsersRequest.UserName = req.UserName
	getUsersRequest.Role = req.Role
	getUsersResponse, err := service.GetUsers(getUsersRequest)
	if err != nil {
		logger.Error(err)
		return err
	}
	rsp.PageNum = getUsersResponse.PageNum
	rsp.PageSize = getUsersResponse.PageSize
	rsp.Count = getUsersResponse.Count
	for _, userInfo := range getUsersResponse.SysUsers {
		var userInfo1 user.GetUserResponse
		userInfo1.ID = userInfo.ID
		userInfo1.UserName = userInfo.UserName
		userInfo1.Email = userInfo.Email
		userInfo1.Phone = userInfo.Phone
		userInfo1.Role = userInfo.Role
		rsp.UsersInfo = append(rsp.UsersInfo, &userInfo1)
	}

	return nil
}

func (e *User) GetUser(ctx context.Context, req *pb.GetUserRequest, rsp *pb.GetUserResponse) error {
	logger.Infof("Received User.GetUser request: %v", req)

	userInfo, err := service.GetUser(req.ID)
	if err != nil {
		logger.Error(err)
		return err
	}
	rsp.ID = userInfo.ID
	rsp.UserName = userInfo.UserName
	rsp.Email = userInfo.Email
	rsp.Phone = userInfo.Phone
	rsp.Role = userInfo.Role
	return nil
}

func (e *User) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest, rsp *pb.DeleteUserResponse) error {
	logger.Infof("Received User.GetUser request: %v", req)

	err := service.DeleteUser(req.ID)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (e *User) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest, rsp *pb.UpdateUserResponse) error {
	logger.Infof("Received User.GetUser request: %v", req)
	var updateUserInfoRequire service.UpdateUsersInfoRequire
	updateUserInfoRequire.UserId = req.UserId
	updateUserInfoRequire.Email = req.Email
	updateUserInfoRequire.Phone = req.Phone
	updateUserInfoRequire.Role = req.Role
	err := service.UpdateUsersInfo(updateUserInfoRequire)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
