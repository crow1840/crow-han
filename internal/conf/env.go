package conf

const (
	ShortTokenKeyPrefix = "crow:token:"
	TokenTimeOut        = 3600 * 24
)

var MyEnvs SysEnvs

type SysEnvs struct {
	MysqlConf
	RedisConf
	Logger
	Consul
}

type MysqlConf struct {
	MysqlIp       string `default:"10.10.239.136" split_words:"true"`
	MysqlPort     string `default:"3306" split_words:"true"`
	DbName        string `default:"crow" split_words:"true"`
	MysqlUser     string `default:"root" split_words:"true"`
	MysqlPassword string `default:"1234" split_words:"true"`
}

//环境变量定义示例：CROW_MYSQL_IP=127.0.0.1，否则使用默认值

type RedisConf struct {
	RidesIp       string `default:"10.10.239.136" split_words:"true"`
	RidesPort     string `default:"6379" split_words:"true"`
	RidesDbNum    int    `default:"1" split_words:"true"`
	RidesPassword string `default:"" split_words:"true"`
}

type Logger struct {
	LogDir       string `default:"logs" split_words:"true" yaml:"dir"`
	LogLevel     string `default:"INFO" split_words:"true" yaml:"level"`
	LogKeepHours uint   `default:"24" split_words:"true" yaml:"keepHours"`
}

type Consul struct {
	ConsulAddr string `default:"10.10.239.32:8500" split_words:"true"`
}
