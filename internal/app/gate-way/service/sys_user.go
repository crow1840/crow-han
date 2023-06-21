package service

import (
	"context"
	"crow-han/internal/pkg"
	userProto "crow-han/proto/user"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/logger"
)

// UpdateUserSelfInfo 用户修改自己信息
// @Summary 用户修改自己信息
// @Description 用户修改自己信息
// @Tags user
// @Accept  application/json
// @Product application/json
// @Param data body UpdateUserSelfInfoRequest  true "选项可选"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/user/info [put]
// @Security Bearer
func UpdateUserSelfInfo(c *gin.Context) {
	var updateUserInfoRequest UpdateUserSelfInfoRequest
	err := c.Bind(&updateUserInfoRequest)
	if err != nil {
		SetErr(c, 500, err, err.Error())
		return
	}
	userIdAny, _ := c.Get("user_id")

	//调用user服务创建用户
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.UpdateUserSelfInfo(context.Background(), &userProto.UpdateUserSelfInfoRequest{
		UserId: userIdAny.(int64),
		Phone:  updateUserInfoRequest.Phone,
		Email:  updateUserInfoRequest.Email,
	})
	if err != nil {
		SetErr(c, 500, err, err.Error())
		return
	}
	logger.Infof("%+v", resp)
	SetOK(c, "用户信息修改成功")
}

// UpdateUserSelfPassword 用户修改密码
// @Summary 用户修改密码
// @Description 用户修改自己的密码
// @Tags user
// @Accept  application/json
// @Product application/json
// @Param data body UpdateUserPasswordRequest  true "用户的新密码和旧密码"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/user/password [put]
// @Security Bearer
func UpdateUserSelfPassword(c *gin.Context) {
	var updateUserPasswordRequest UpdateUserPasswordRequest

	err := c.Bind(&updateUserPasswordRequest)
	if err != nil {
		SetErr(c, 500, err, err.Error())
		return
	}
	userIdAny, _ := c.Get("user_id")
	userName, _ := c.Get("user_name")
	logger.Infof("userId : %+v, userName: %+v", userIdAny, userName)
	//旧密码认证
	passwordMd5 := pkg.EncryptWithMd5(updateUserPasswordRequest.Password)
	logger.Info("PasswordMd5 :", passwordMd5)

	//调用user
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.UpdateUserSelfPassword(context.Background(), &userProto.UpdateUserSelfPasswordRequest{
		UserId:      userIdAny.(int64),
		UserName:    userName.(string),
		Password:    updateUserPasswordRequest.Password,
		NewPassword: updateUserPasswordRequest.NewPassword,
	})
	if err != nil {
		SetErr(c, 500, err, err.Error())
		return
	}
	logger.Infof("%+v", resp)

	SetOK(c, "密码修改成功")
}

// GetUserSelfInfo 用户获取本人信息
// @Summary 用户获取本人信息
// @Description 用户获取本人信息
// @Tags user
// @Success 200 {object} response.Response{data=user.GetUserSelfInfoResponse} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/user/info [get]
// @Security Bearer
func GetUserSelfInfo(c *gin.Context) {

	userNameAny, _ := c.Get("user_name")
	//if b == false {
	//	err := errors.New("用户认证失败")
	//	SetErr(c, 500, err, err.Error())
	//}

	//调用user
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.GetUserSelfInfo(context.Background(), &userProto.GetUserSelfInfoRequest{
		UserName: userNameAny.(string),
	})
	if err != nil {
		SetErr(c, 500, err, err.Error())
		return
	}

	logger.Info(resp)
	SetOK(c, resp)

}
