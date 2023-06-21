package service

import (
	"context"
	"crow-han/internal/conf"
	authProto "crow-han/proto/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/logger"
	"strings"
)

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录，并获取token
// @Tags system
// @Accept  application/json
// @Product application/json
// @Param data body LoginRequest  true "用户名，用户密码"
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/login [post]
// @Security Bearer
func Login(c *gin.Context) {

	var loginRequest LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		SetErr(c, 500, err, err.Error())
		return
	}
	logger.Infof("%+v", loginRequest)

	//調用auth
	authSrv := authProto.NewAuthService("auth", srv.Client())
	resp, err := authSrv.Login(context.Background(), &authProto.LoginRequest{
		UserName: loginRequest.UserName,
		Password: loginRequest.Password,
	})
	if err != nil {
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}
	logger.Infof("request : %+v", resp)
	SetOK(c, resp.Token)
}

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出
// @Tags system
// @Accept  application/json
// @Product application/json
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/logout [post]
// @Security Bearer
func Logout(c *gin.Context) {

	//获取token
	authorization := c.Request.Header.Get("Authorization")
	logger.Infof("authorization：%+v", authorization)

	s2 := strings.SplitN(authorization, " ", 2)
	logger.Infof("s2：%+v", s2)
	strToken := s2[1]
	if strToken == "" {
		err := errors.New("获取token失败")
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}
	logger.Infof("strToken：%s", strToken)

	//删除redis中的token
	redisValue := conf.Rdb.Del(conf.RedisCtx, conf.ShortTokenKeyPrefix+strToken)

	SetOK(c, redisValue)
}

// RefreshToken 刷新token
// @Summary 刷新token
// @Description 刷新用户token
// @Tags system
// @Accept  application/json
// @Product application/json
// @Success 200 {object} response.Response{data=string} "{"requestId": "string","code": 200,"msg": "ok","data": [...]}"
// @Failure 500 {object} response.Response{msg=string} "{"requestId": "string","code": 500,"msg": "string","status": "error","data": null}"
// @Router /api/v1/login/refresh [post]
// @Security Bearer
func RefreshToken(c *gin.Context) {
	//获取token
	authorization := c.Request.Header.Get("Authorization")
	logger.Infof("authorization：%+v", authorization)

	s2 := strings.SplitN(authorization, " ", 2)
	logger.Infof("s2：%+v", s2)
	strToken := s2[1]
	if strToken == "" {
		err := errors.New("获取token失败")
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}
	logger.Infof("strToken：%s", strToken)

	//调用auth
	authSrv := authProto.NewAuthService("auth", srv.Client())
	resp, err := authSrv.RefreshToken(context.Background(), &authProto.RefreshTokenRequest{
		Token: strToken,
	})
	if err != nil {
		logger.Error(err)
		SetErr(c, 500, err, err.Error())
		return
	}
	SetOK(c, resp.NewToken)
}

func admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		authorization := c.Request.Header.Get("Authorization")
		var err error
		s2 := strings.SplitN(authorization, " ", 2)
		//fmt.Printf("===================%+v", s2)
		strToken := s2[1]
		//fmt.Printf("===================%+v", strToken)
		if strToken == "" {
			err = errors.New("获取token失败")
			SetErr(c, 500, err, err.Error())
		}

		authSrv := authProto.NewAuthService("auth", srv.Client())
		resp, err := authSrv.Verify(context.Background(), &authProto.VerifyRequest{
			Token: strToken,
		})
		logger.Infof("%+v", resp)
		if err != nil {
			c.Abort()
		}
		if resp.Result == false {
			err = errors.New("身份认证失败")
			SetErr(c, 500, err, err.Error())
			c.Abort()
		}

		logger.Info(resp.UserRole)
		if resp.UserRole != "admin" {
			err := errors.New("没有管理员权限")
			SetErr(c, 500, err, err.Error())
			c.Abort()
		}
		c.Next()
	}
}

func user() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		var err error
		s2 := strings.SplitN(authorization, " ", 2)
		//fmt.Printf("===================%+v", s2)
		strToken := s2[1]
		//fmt.Printf("===================%+v", strToken)
		if strToken == "" {
			err = errors.New("获取token失败")
			SetErr(c, 500, err, err.Error())
		}
		//调用auth验证
		authSrv := authProto.NewAuthService("auth", srv.Client())
		resp, err := authSrv.Verify(context.Background(), &authProto.VerifyRequest{
			Token: strToken,
		})
		logger.Infof("%+v", resp)
		if err != nil {
			c.Abort()
		}
		if resp.Result == false {
			err = errors.New("身份认证失败")
			SetErr(c, 500, err, err.Error())
			c.Abort()
		}

		c.Set("user_role", resp.UserRole)
		c.Set("user_name", resp.UserName)
		c.Set("user_id", resp.UserId)
		c.Next()
	}

}
