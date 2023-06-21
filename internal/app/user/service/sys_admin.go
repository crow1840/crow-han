package service

import (
	"crow-han/internal/conf"
	"crow-han/internal/models"
	"crow-han/internal/pkg"
	"errors"
	"github.com/toolkits/pkg/logger"
	"regexp"
)

//调用

func CreateUser(createUserRequire CreateUserRequire) (userId int64, err error) {

	//输入检查
	if createUserRequire.UserName == "" {
		err = errors.New("用户名不能为空")
		logger.Error(err)
		return userId, err
	}
	if createUserRequire.Role == "" {
		err = errors.New("用户角色不能为空")
		logger.Error(err)
		return userId, err
	}
	b, err := checkPassword(createUserRequire.Password)
	if b == false {
		logger.Error(err)
		return userId, err
	}

	//查重
	checkResult, err := models.CheckUserExists(createUserRequire.UserName, conf.MyTx)
	if checkResult == true {
		err = errors.New("用户名已经存在")
		logger.Error(err)
		return userId, err
	} else {
		logger.Info("用户名查重通过")
	}

	//创建用户
	var sysUser models.SysUsers
	sysUser.UserName = createUserRequire.UserName
	sysUser.Role = createUserRequire.Role
	sysUser.Phone = createUserRequire.Phone
	sysUser.Email = createUserRequire.Email
	sysUser.Password = pkg.EncryptWithMd5(createUserRequire.Password)

	userId, err = models.CreateUser(sysUser, conf.MyTx)
	if err != nil {
		logger.Error(err)
	}
	return userId, nil

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

//调用

func ResetUserPassword(resetUserPasswordRequire ResetUserPasswordRequire) (err error) {

	//输入检查
	if &resetUserPasswordRequire.UserId == nil {
		err = errors.New("用户id不能为空")
		logger.Error(err)
		return err
	}

	b, err := checkPassword(resetUserPasswordRequire.NewPassword)
	if b == false {
		logger.Error(err)
		return err
	}

	//检查用户是否存在
	updateUser, err := models.GetUserById(resetUserPasswordRequire.UserId, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	if len(updateUser.UserName) == 0 {
		err = errors.New("用户不存在")
		logger.Error(err)
		return err
	}

	//修改密码
	var sysUsers models.SysUsers
	sysUsers.ID = resetUserPasswordRequire.UserId
	sysUsers.Password = pkg.EncryptWithMd5(resetUserPasswordRequire.NewPassword)
	logger.Info(sysUsers)
	err = models.UpdateUserInfoByID(sysUsers, conf.MyTx)
	return nil
}

//调用

func GetUsers(getUsersRequest GetUsersRequire) (getUsersResponse GetUsersResponse, err error) {

	//查询
	var sysUser models.SysUsers
	sysUser.UserName = getUsersRequest.UserName
	sysUser.Role = getUsersRequest.Role
	sysUsers, count, err := models.GetUsersPage(getUsersRequest.PageSize, getUsersRequest.PageNum, sysUser, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return getUsersResponse, err
	}

	//组装返回结果
	getUsersResponse.SysUsers = sysUsers
	getUsersResponse.PageSize = getUsersRequest.PageSize
	getUsersResponse.PageNum = getUsersRequest.PageNum
	getUsersResponse.Count = count
	return getUsersResponse, nil
}

//调用

func GetUser(userID int64) (sysUser models.SysUsers, err error) {

	//查询
	sysUser, err = models.GetUserById(userID, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return sysUser, err
	}
	return sysUser, nil
}

//调用

func DeleteUser(userID int64) (err error) {

	//检查用户是否存在
	updateUser, err := models.GetUserById(userID, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	if len(updateUser.UserName) == 0 {
		err = errors.New("用户不存在")
		logger.Error(err)
		return err
	}

	//删除用户
	err = models.DeleteUserById(userID, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

//调用

func UpdateUsersInfo(updateUserInfoRequire UpdateUsersInfoRequire) (err error) {
	//检查用户是否存在
	updateUser, err := models.GetUserById(updateUserInfoRequire.UserId, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	if len(updateUser.UserName) == 0 {
		err = errors.New("用户不存在")
		logger.Error(err)
		return err
	}
	//更新数据库
	var sysUsers models.SysUsers
	sysUsers.ID = updateUserInfoRequire.UserId
	sysUsers.Email = updateUserInfoRequire.Email
	sysUsers.Phone = updateUserInfoRequire.Phone
	sysUsers.Role = updateUserInfoRequire.Role
	logger.Info(sysUsers)
	err = models.UpdateUserInfoByID(sysUsers, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
