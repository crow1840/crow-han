package conf

import (
	"database/sql"
	"fmt"
	"github.com/toolkits/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type MyConnect struct {
	Tx    *gorm.DB
	SqlDB *sql.DB
	Err   error
}

var MyTx *gorm.DB

func (myTx *MyConnect) ConnectMysql() {
	var db *gorm.DB
	var sqlDB *sql.DB
	var err error
	dsn := MyEnvs.MysqlUser + ":" + MyEnvs.MysqlPassword + "@tcp(" + MyEnvs.MysqlIp + ":" + MyEnvs.MysqlPort + ")/" + MyEnvs.DbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ = db.DB()
	if err != nil {
		fmt.Printf(err.Error())
		defer sqlDB.Close()
	} else {
		logger.Info("数据库链接成功")
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
	MyTx = db
}
