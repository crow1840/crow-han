package service

import "github.com/dgrijalva/jwt-go"

//type ResponsesLogin struct {
//	err   error  `json:"err"`
//	token string `json:"token"`
//}

type LoginInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID      int      `json:"user_id"`
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	FullName    string   `json:"full_name"`
	Permissions []string `json:"permissions"`
}

var (
	Secret     = "LiuBei"  // 加盐
	ExpireTime = 3600 * 24 // token有效期
)
