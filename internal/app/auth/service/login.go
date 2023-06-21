package service

import (
	"crow-han/internal/conf"
	"crow-han/internal/models"
	"crow-han/internal/pkg"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/toolkits/pkg/logger"
	"time"
)

func Login(loginRequest LoginInfo) (signedToken string, err error) {

	passwordMd5 := pkg.EncryptWithMd5(loginRequest.Password)
	logger.Info("PasswordMd5 :", passwordMd5)

	result, err := models.CheckUser(loginRequest.UserName, passwordMd5, conf.MyTx)
	if err != nil {
		logger.Error(err)
		return signedToken, err
	}
	logger.Infof("%s 用户比对结果: %t", loginRequest.UserName, result)

	if result == false {
		err = errors.New("用户名或密码错误")
		logger.Error(err)
		return signedToken, err
	}

	claims := &JWTClaims{
		UserID:      1,
		Username:    loginRequest.UserName,
		Password:    loginRequest.Password,
		FullName:    loginRequest.UserName,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err = getToken(claims)

	if err != nil {
		return signedToken, err
	}
	//fmt.Printf(signedToken)

	//存入redis
	redisKey := fmt.Sprintf("%s%s", conf.ShortTokenKeyPrefix, signedToken)
	//fmt.Println("==============", key)
	err = conf.Rdb.Set(conf.RedisCtx, redisKey, claims.Username, time.Second*conf.TokenTimeOut).Err()
	if err != nil {
		logger.Info(err)
	}
	return signedToken, nil
}

//身份验证

func Verify(strToken string) (result bool, userId int, userName string, userRole string, err error) {

	//验证token，并获取用户名
	claim, err := verifyAction(strToken)
	if err != nil {
		result = false
		return result, 0, "", "", nil
	}
	result = true
	userName = claim.Username
	logger.Infof("claim : %+v", claim)
	userId = claim.UserID
	//获取用户角色
	userRole, err = models.GetRoleByUserName(userName, conf.MyTx)

	//对比redis中的token
	redisValue := conf.Rdb.Get(conf.RedisCtx, conf.ShortTokenKeyPrefix+strToken)
	if redisValue.Val() != userName {
		err = errors.New("用户token验证失败：002")
		return false, userId, userName, userRole, err
	}
	logger.Infof("redis中的用户名：%s，token的用户名：%s", redisValue.Val(), claim.Username)

	return result, userId, userName, userRole, nil

}

//刷新token

func RefreshToken(token string) (newToken string, err error) {

	//验证token，并获取用户名
	claims, err := verifyAction(token)
	if err != nil {
		return newToken, err
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	newToken, err = getToken(claims)
	if err != nil {
		return newToken, err
	}
	//存入redis
	redisKey := fmt.Sprintf("%s%s", conf.ShortTokenKeyPrefix, newToken)
	//fmt.Println("==============", key)
	err = conf.Rdb.Set(conf.RedisCtx, redisKey, claims.Username, time.Second*conf.TokenTimeOut).Err()
	if err != nil {
		logger.Info(err)
	}

	//删除旧token
	err = conf.Rdb.Del(conf.RedisCtx, conf.ShortTokenKeyPrefix+token).Err()
	if err != nil {
		logger.Error(err)
	}
	return newToken, nil
}

// 解析token，获取用户信息
func verifyAction(strToken string) (claims *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, err
	}
	if err = token.Claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}

// 生成token

func getToken(claims *JWTClaims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
