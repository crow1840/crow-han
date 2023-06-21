package service

import (
	"crow-han/internal/conf"
	"crow-han/internal/models"
	"crow-han/internal/pkg"
	"errors"
	"github.com/toolkits/pkg/logger"
)

// UpdateUserInfo 用户修改自己信息

func UpdateUserSelfInfo(updateUserInfoRequire UpdateUserSelfInfoRequire) (err error) {

	//更新数据库
	var sysUsers models.SysUsers
	sysUsers.ID = updateUserInfoRequire.UserID
	sysUsers.Email = updateUserInfoRequire.Email
	sysUsers.Phone = updateUserInfoRequire.Phone
	logger.Info(sysUsers)
	err = models.UpdateUserInfoByID(sysUsers, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

// UpdateUserPassword 用户修改密码

func UpdateUserSelfPassword(updateUserPasswordRequire UpdateUserSelfPasswordRequire) (err error) {
	logger.Info(updateUserPasswordRequire)

	//旧密码认证
	passwordMd5 := pkg.EncryptWithMd5(updateUserPasswordRequire.Password)
	logger.Info("PasswordMd5 :", passwordMd5)

	result, err := models.CheckUser(updateUserPasswordRequire.UserName, passwordMd5, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Infof("%s 用户比对结果: %t", updateUserPasswordRequire.UserName, result)

	if result == false {
		err = errors.New("密码错误")
		logger.Error(err)
		return err
	}

	//修改密码
	var sysUsers models.SysUsers
	sysUsers.UserName = updateUserPasswordRequire.UserName
	sysUsers.Password = pkg.EncryptWithMd5(updateUserPasswordRequire.NewPassword)
	logger.Info(sysUsers)
	err = models.UpdateUserInfoByUserName(sysUsers, conf.MyTx)
	return nil
}

// GetUserInfo 用户获取本人信息

func GetUserSelfInfo(userName string) (userInfo models.SysUsers, err error) {

	//查询信息
	userInfo, err = models.GetUserByUserName(userName, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return userInfo, err
	}
	logger.Info(userInfo)
	return userInfo, nil
}
