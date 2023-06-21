package cache

import (
	"crow-han/internal/conf"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/logger"
	"time"
)

type UserName struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Age      int64  `json:"age"`
}

func example(c *gin.Context) {

	data := &UserName{
		UserId:   1,
		UserName: "LiuBei",
		Age:      30,
	}

	//结构体转为json字串的[]byte
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	//写入
	err = conf.Rdb.Set(conf.RedisCtx, "project:users:{user_name}:struct", string(b), time.Minute*2).Err()
	if err != nil {
		fmt.Println("err: ", err)
	}

	//查找
	result := conf.Rdb.Get(conf.RedisCtx, "project:users:{user_name}:struct")
	fmt.Println(result.Val())
}

func CheckRides() {
	err := conf.Rdb.Set(conf.RedisCtx, "crow-han:init", "OK", time.Second*20).Err()
	if err != nil {
		logger.Info(err)
	}

	result := conf.Rdb.Get(conf.RedisCtx, "crow-han:init")
	logger.Info(result.Val())
	if result.Val() == "OK" {
		logger.Info("redis 链接成功")
	} else {
		logger.Error("reids 链接失败")
	}
	fmt.Println(result.Val())
}
