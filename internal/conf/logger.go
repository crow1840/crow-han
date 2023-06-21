package conf

import (
	"fmt"
	"github.com/toolkits/pkg/logger"
	"os"
)

func LoggerInit() {
	logger.Infof("=================", MyEnvs.LogDir)
	lb, err := logger.NewFileBackend(MyEnvs.LogDir)
	if err != nil {
		fmt.Println("cannot init logger:", err)
		os.Exit(1)
	}
	//来开启按小时滚动
	lb.SetRotateByHour(true)
	//保留24小时的log
	lb.SetKeepHours(MyEnvs.LogKeepHours)
	logger.SetLogging(MyEnvs.LogLevel, lb)
}
