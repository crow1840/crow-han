# 1.数据库链接
### 1.1 使用sdk：gorm
### 1.2 配置
- 默认配置

`/src/cofig/evn.go` 文件中，修改default中内容

```go
type MysqlConf struct {
	MysqlIp       string `default:"10.10.xxx.xxx" split_words:"true"`
	MysqlPort     string `default:"3306" split_words:"true"`
	DbName        string `default:"crow" split_words:"true"`
	MysqlUser     string `default:"root" split_words:"true"`
	MysqlPassword string `default:"xxxxx" split_words:"true"`
}
```
- 用环境变量指定
```shell
export CROW_MYSQL_IP="10.10.xxx.xxx"
export CROW_MYSQL_PORT="3306"
export CROW_DB_NAME="crow"
export CROW_MYSQL_USER="root"
export CROW_MYSQL_PASSWORD="xxxxxxx"
```
- docker-compose 中指定
```yml
    environment:
      - CROW_MYSQL_IP=10.10.xxx.xxx
      - CROW_MYSQL_PORT=3306
      - CROW_MYSQL_DB_NAME=crow
      - CROW_MYSQL_USER=root
      - CROW_MYSQL_PASSWORD=xxxxxx
```

# 2. redis 链接
## 2.1 使用sdk：
"github.com/go-redis/redis/v8"

## 2.2 配置
- 默认配置

`/src/cofig/evn.go` 文件中，修改default中内容

```go
type RedisConf struct {
    RidesIp       string `default:"10.10.239.136" split_words:"true"`
    RidesPort     string `default:"6379" split_words:"true"`
    RidesDbNum    int    `default:"1" split_words:"true"`
    RidesPassword string `default:"" split_words:"true"`
}
```
- 用环境变量指定
```shell
export CROW_RIDES_IP="10.10.xxx.xxx"
export CROW_RIDES_PORT="6379"
export CROW_RIDES_DB_NUM=2
export CROW_RIDES_PASSWORD="xxxxxxx"
```
- docker-compose 中指定
```yml
    environment:
      - CROW_RIDES_IP=10.10.xxx.xxx
      - CROW_RIDES_PORT=6379
      - CROW_RIDES_DB_NUM=2
      - CROW_RIDES_PASSWORD="xxxxxxxx"
```

# 3. 日志
## 3.1 日志sdk
"github.com/toolkits/pkg/logger"

## 3.2 配置
- 默认配置
```go
type Logger struct {
	LogDir       string `default:"logs" split_words:"true" yaml:"dir"`
	LogLevel     string `default:"INFO" split_words:"true" yaml:"level"`
	LogKeepHours uint   `default:"24" split_words:"true" yaml:"keepHours"`
}
```
- 用环境变量指定
```shell
export CROW_LOG_DIR="logs"
export CROW_LOG_LEVEL="ERROR"
export CROW_KEEP_HOURS=24*7
```
- docker-compose 中指定
```yml
    environment:
      - CROW_LOG_DIR=logs
      - CROW_RIDES_PORT=ERROR
      - CROW_KEEP_HOURS=144
```

# 4. protoc
```shell
 protoc --proto_path=. --micro_out=. --go_out=:. proto/auth/auth.proto
 protoc --proto_path=. --micro_out=. --go_out=:. proto/gate-way/gate-way.proto
 protoc --proto_path=. --micro_out=. --go_out=:. proto/user/user.proto
```

# 5. swagger
- 编译命令
```shell
cd internal/app/gate-way/service
swag init --parseDependency --parseInternal --parseDepth 2 -g .\router.go
```

---
# 相关文档
- 快速启动 
https://blog.csdn.net/xingzuo_1840/article/details/130618097

- 开发实战
https://blog.csdn.net/xingzuo_1840/article/details/130679235