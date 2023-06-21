package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
)

type OK struct {
	Code   int64       `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Status string      `json:"status"`
}

func SetOK(c *gin.Context, data interface{}) {
	c.Set("status", "ok")
	response.OK(c, data, "ok")
}

func SetErr(c *gin.Context, code int, err error, msg string) {
	response.Error(c, code, err, msg)
}
