package service

import (
	"context"
	"crow-han/internal/pkg"
	userProto "crow-han/proto/user"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/logger"
	"regexp"
)

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags admin
// @Accept  application/json
// @Product application/json
// @Param data body CreateUserRequest  true "填写用户信息"
// @Success 200 {object} response.Response{data=int64} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/user [post]
// @Security Bearer
func CreateUser(c *gin.Context) {

	var createUserRequest CreateUserRequest
	err := c.Bind(&createUserRequest)
	if err != nil {
		SetErr(c, 500, err, err.Error())
		return
	}

	//输入检查
	if createUserRequest.UserName == "" {
		err = errors.New("用户名不能为空")
		SetErr(c, 500, err, "用户名不能为空")
		//c.JSON(403, gin.H{"err": "用户名不能为空"})
		return
	}
	if createUserRequest.Role == "" {
		err = errors.New("用户角色不能为空")
		SetErr(c, 500, err, "用户角色不能为空")
		return
	}
	b, err := checkPassword(createUserRequest.Password)
	if b == false {
		SetErr(c, 500, err, err.Error())
		return
	}

	//调用user服务创建用户
	//userSrv := *UserSrv
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.CreateUser(context.Background(), &userProto.CreateUserRequest{
		UserName: createUserRequest.UserName,
		Password: createUserRequest.Password,
		Email:    createUserRequest.Email,
		Phone:    createUserRequest.Phone,
		Role:     createUserRequest.Role,
	})
	if err != nil {
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}
	logger.Infof("%+v", resp)

	SetOK(c, resp.UserId)
}

func checkPassword(password string) (b bool, err error) {
	//校验密码长度
	if len(password) > 18 || len(password) < 6 {
		return false, errors.New("密码必须是8~16位")
	} else {
		logger.Info(len(password), "长度检查通过")
	}

	//校验纯英文
	b, err = regexp.MatchString("^([A-z]+)$", password)
	if err != nil {
		logger.Error(err)
	}
	if b == true {
		return false, errors.New("密码不能是纯英文")
	} else {
		logger.Info("纯英文检查通过")
	}

	//校验纯数字
	b, err = regexp.MatchString("^([0-9]+)$", password)
	if err != nil {
		logger.Error(err)
	}
	if b == true {
		return false, errors.New("密码不能是纯数字")
	} else {
		logger.Info("纯数字检查通过")
	}
	return true, nil
}

// ResetUserPassword admin重置用户密码
// @Summary admin重置用户密码
// @Description admin重置用户密码
// @Tags admin
// @Accept  application/json
// @Product application/json
// @Param data body ResetUserPasswordRequest  true "填写用户Id和xin密码"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/user-password [put]
// @Security Bearer
func ResetUserPassword(c *gin.Context) {
	var resetUserPasswordRequest ResetUserPasswordRequest
	err := c.Bind(&resetUserPasswordRequest)
	if err != nil {
		SetErr(c, 403, err, err.Error())
		return
	}

	//输入检查
	if &resetUserPasswordRequest.UserId == nil {
		SetErr(c, 403, errors.New("用户id不能为空"), "用户id不能为空")
	}

	b, err := checkPassword(resetUserPasswordRequest.NewPassword)
	if b == false {
		SetErr(c, 500, err, err.Error())
		return
	}

	//调用user
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.ResetUserPassword(context.Background(), &userProto.ResetUserUserRequest{
		UserId:      resetUserPasswordRequest.UserId,
		NewPassword: resetUserPasswordRequest.NewPassword,
	})
	logger.Infof("%+v", resp)
	if err != nil {
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}

	SetOK(c, "密码重置成功")
}

// GetUsers 获取用户信息
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags admin
// @Param role query string false "用户角色"
// @Param user_name query string false "用户名（可模糊查询）"
// @Param page_num query string false "页数"
// @Param page_size query string false "每页行数"
// @Success 200 {object} response.Response{data=user.GetUsersResponse} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/users [get]
// @Security Bearer
func GetUsers(c *gin.Context) {

	pageNumString := c.Query("page_num")
	pageNum := pkg.StringToInt64(pageNumString)
	pageSizeString := c.Query("page_size")
	pageSize := pkg.StringToInt64(pageSizeString)
	userName := c.Query("user_name")
	role := c.Query("role")

	//调用user
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.GetUsers(context.Background(), &userProto.GetUsersRequest{
		PageNum:  pageNum,
		PageSize: pageSize,
		UserName: userName,
		Role:     role,
	})
	logger.Infof("==================%+v===============", resp)
	if err != nil {
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}

	SetOK(c, resp)
}

// GetUser 获取指定用户信息
// @Summary 获取指定用户信息
// @Description 获取指定用户信息
// @Tags admin
// @Param uuid path string true "用户id"
// @Success 200 {object} response.Response{data=user.GetUserResponse} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/users/{uuid} [get]
// @Security Bearer
func GetUser(c *gin.Context) {
	uuid := c.Param("uuid")

	////查询
	//var sysUser models.SysUsers
	userId := pkg.StringToInt64(uuid)

	//调用user
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.GetUser(context.Background(), &userProto.GetUserRequest{
		ID: userId,
	})
	logger.Infof("==================%+v===============", resp)
	if err != nil {
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}

	SetOK(c, resp)

}

// DeleteUser 删除指定用户
// @Summary 删除指定用户
// @Description 删除指定用户
// @Tags admin
// @Param uuid path string true "用户id"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/users/{uuid} [delete]
// @Security Bearer
func DeleteUser(c *gin.Context) {
	uuid := c.Param("uuid")

	////查询
	//var sysUser models.SysUsers
	userId := pkg.StringToInt64(uuid)

	//调用user
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.DeleteUser(context.Background(), &userProto.DeleteUserRequest{
		ID: userId,
	})
	logger.Infof("==================%+v===============", resp)
	if err != nil {
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}

	SetOK(c, "删除成功")

}

// UpdateUserInfo 修改用户信息
// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags admin
// @Accept  application/json
// @Product application/json
// @Param data body UpdateUsersInfoRequest  true "user_id必须，其他可选"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/admin/users [put]
// @Security Bearer
func UpdateUserInfo(c *gin.Context) {
	var updateUserInfoRequest UpdateUsersInfoRequest
	err := c.Bind(&updateUserInfoRequest)
	if err != nil {
		SetErr(c, 500, err, err.Error())
		return
	}
	logger.Info(updateUserInfoRequest)
	//调用user
	userSrv := userProto.NewUserService("user", srv.Client())
	resp, err := userSrv.UpdateUser(context.Background(), &userProto.UpdateUserRequest{
		UserId: updateUserInfoRequest.UserId,
		Email:  updateUserInfoRequest.Email,
		Phone:  updateUserInfoRequest.Phone,
		Role:   updateUserInfoRequest.Role,
	})
	logger.Info(resp)
	if err != nil {
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}
	SetOK(c, "用户信息修改成功")

}
