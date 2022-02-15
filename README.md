# Ginco

## 介绍
Ginco是一个Golang框架，基于gin框架和cobra CLI库实现，开箱即用。大部分服务基于契约，均可替换。

##### 目前实现的服务：

| 服务 | 别名 | 备注 |
| --- | --- | --- |
| config | - | 配置服务，基于contract.Config契约，底层使用viper.Viper |
| console | cmd | 命令行服务，使用cobra.Command |
| http | server、router | HTTP服务，底层使用gin.Engine |
| logger | log | 日志服务，基于contract.Logger契约，底层使用zap.Logger |
| redis | - | Redis服务，基于contract.Redis契约，底层使用redis.UniversalClient |
| database | db | 数据库服务，基于contract.Database契约，底层使用gorm.DB |
| cache | - | 缓存服务，基于contract.Cache契约，支持redis、memory、database三种驱动 |
| validate | validator | 数据验证服务，使用validator.Validate |

> 其他服务待实现，当然你也可以自己直接注册服务（Set），或者实现contract.Provider接口，然后绑定（Bind）即可


## 安装
直接下载即可
```shell
git clone https://github.com/zhaoyang1214/ginco.git
```

## 快速使用
1. 进入项目目录，运行`go run main.go`即可看到可执行的命令
```shell
Usage:
   [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  restart     Restart http server
  start       Start http server (alias: server)
  stop        Stop http server
  version     Get Application version

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```

2. 运行`go run main.go start`即可拉起HTTP服务， 请求`http://127.0.0.1:8080`即可看到输出
`Hello Ginco v0.0.0`  
要想进入daemon模式，可运行`go run main.go start -d`（**该模式只能运行在UNIX-based OS**）  
使用 `go run main.go start -h`即可看到支持的参数
```shell
Run attaches the router to a http.Server and starts listening and serving HTTP requests.

Usage:
   start [flags]

Aliases:
  start, server

Flags:
  -d, --daemon     Start http.Server daemon
  -h, --help       help for start
  -p, --port int   Listening and serving HTTP port (default 8080)
```
可以看到支持端口参数（-p），默认使用配置文件`config/http.yaml`中的`http.port`配置

3. 同时还支持`stop`和`restart`，用来停止和重启HTTP服务

## 目录结构
```
├── app
│   ├── console     #用户命令
│   │   ├── command
│   │   │   └── version
│   │   │       └── version.go
│   │   └── kernel.go
│   ├── http        #http相关业务服务
│   │   ├── controller
│   │   │   └── index.go
│   │   └── middleware  #gin中间件
│   └── providers       #自定义服务提供者
│       └── kernel.go
├── bootstrap       #初始化应用
│   ├── app.go
│   └── kernel.go
├── config          #配置文件目录
│   ├── app.yaml
│   ├── http.yaml
│   └── logger.yaml
├── framework       #框架核心目录
├── go.mod
├── go.sum
├── main.go         
├── README.md
├── router          #注册路由
│   ├── api.go
│   └── kernel.go
└── runtime         #运行时产生的目录
```

## 配置信息
配置文件放在`config`目录下，支持`json, toml, yaml, yml, properties, props, prop, hcl, tfvars, dotenv, env, ini`。  

项目根目录下如果存在`.env`文件，则会自动加载到配置中。同时允许配置文件使用变量，例如`${APP_NAME}`，会自动替换成`.env`文件中的`APP_NAME`值

获取配置
```shell
// a contract.Application
a.GetI("config").(contract.Config).GetString("app.name")
```

## 服务容器
全局获取容器
```shell
a := app.Get()
```

注册服务
```shell
// a contract.Application
a.Set("serverName", server)
```

获取服务
```shell
// a contract.Application
configServer,err := a.Get("config")
config := configServer.(contract.Config)
// or
config := a.GetI("config").(contract.Config)
```

绑定Provider
```shell
// a contract.Application
a.Bind("serverName", server)
```

给服务设置别名
```shell
// a contract.Application
a.Alias("config", "conf")
// 获取config服务
conf := a.GetI("conf").(contract.Config)
```

## 服务提供者
实现`contract.Provider`接口，然后绑定到容器`a.Bind("serverName", server)`即可

## 契约
使用契约是为了解耦。目前约定了部分契约，后续还会增加

## 路由
路由使用的是`gin.Engine`，只要在`router.Register`中注册相应的路由即可。
```shell
router := a.GetI("router").(*gin.Engine)
router.GET("/", func (c *gin.Context) {
	c.String(http.StatusOK, "Hello Ginco v"+a.Version()+"\n")
})
```

具体使用请参考https://github.com/gin-gonic/gin

## 中间件
在`router.Register`中注册相应的中间件即可。
```shell
router := a.GetI("router").(*gin.Engine)
router.Use(gin.Logger(), gin.Recovery())
```

## 日志
日志服务，基于contract.Logger契约，底层使用zap.Logger (go.uber.org/zap)。  

支持single、rotation、stderr、stack驱动。stack可以配置多个日志通道。  

`rotation`驱动使用`rotatelogs.RotateLogs`包对日志切割，可以按时间（rotationTime）和日志大小（rotationSize）进行切割

#### 配置项
| 配置项 | 支持的驱动 | 备注 |
| --- | --- | --- |
| level | all | 日志级别 |
| encoding | all | 日志输出格式，支持传统的console和json |
| encoderConfig.messageKey | single、rotation、stderr |  |
| encoderConfig.levelKey | single、rotation、stderr |  |
| encoderConfig.timeKey | single、rotation、stderr |  |
| encoderConfig.nameKey | single、rotation、stderr |  |
| encoderConfig.callerKey | single、rotation、stderr |  |
| encoderConfig.functionKey | single、rotation、stderr |  |
| encoderConfig.stacktraceKey | single、rotation、stderr |  |
| encoderConfig.lineEnding | single、rotation、stderr |  |
| encoderConfig.timeEncoder | single、rotation、stderr |  |
| encoderConfig.callerEncoder | single、rotation、stderr |  |
| encoderConfig.consoleSeparator | single、rotation、stderr |  |
| development | all |  |
| disableCaller | all |  |
| callerSkip | all |  |
| disableStacktrace | all |  |
| sampling.initial | all |  |
| sampling.thereafter | all |  |
| initialFields | all | 预定义一些值写入日志 |
| channels | stack |  |
| path | single、rotation | 目录支持的日期格式：%Y%m%d%H%M%S |
| maxAge | rotation |  |
| rotationTime | rotation |  |
| rotationCount | rotation |  |
| rotationSize | rotation |  |


#### 使用
```
log := a.GetI("log").(contract.Logger)
defer log.Sync()

log.Debug("test debug", map[string]string{"t1":"111"})

log.Info("test info", zap.String("t1", "111"))

log.Log(zap.DPanicLevel, "test log")

log.Error("test error", map[string]string{"t1":"111"})

```


## 命令行
命令行使用cobra.Command，在console.Register注册相应的命令即可。  

运行`go run main.go`即可看到所有支持的命令。
```
Usage:
   [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  restart     Restart http server
  start       Start http server (alias: server)
  stop        Stop http server
  version     Get Application version

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command
```

## Redis
Redis服务，基于contract.Redis契约，底层使用redis.UniversalClient（github.com/go-redis/redis/v8）。  

支持**单机**、**集群**、**哨兵**三种模式，根据配置文件中的配置项决定：

1. 如果配置了`master_name`，则是**哨兵**模式
2. 如果`addrs`配置了两个及以上地址，则是**集群**模式
3. 其他情况，则是**单机**模式

#### 使用：
```
redisClient := a.GetI("redis").(contract.Redis)
ctx := context.Background()

// 默认使用`default`配置
err := redisClient.Set(ctx, "test", "111", 0).Err()
if err != nil {
    panic(err)
}

fmt.Println(redisClient.Get(ctx, "test"))

// 使用其他`redis`配置（需在`config/redis.yaml`文件中配置）
r := redisClient.Connection("other_redis")
err := r.Set(ctx, "test1", "ttt", 60 * time.Second).Err()
if err != nil {
    panic(err)
}

fmt.Println(r.Get(ctx, "test1"))

```

#### 配置项
| 配置项 | 支持的模式 | 备注 |
| --- | --- | --- |
| addrs | all | host:port |
| db | 单机、哨兵 |  |
| password | all |  |
| username | all | Redis >= 6.0版本支持 |
| max_retries | all | Maximum number of retries before giving up.Default is 3 retries; -1 (not 0) disables retries. |
| min_retry_backoff | all | Minimum backoff between each retry.Default is 8 milliseconds; -1 disables backoff. |
| max_retry_backoff | all | Maximum backoff between each retry.Default is 512 milliseconds; -1 disables backoff. |
| dial_timeout | all | Dial timeout for establishing new connections.Default is 5 seconds. |
| read_timeout | all | Timeout for socket reads. If reached, commands will fail with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.Default is 3 seconds. |
| write_timeout | all | Timeout for socket writes. If reached, commands will fail with a timeout instead of blocking.Default is ReadTimeout. |
| pool_fifo | all | Type of connection pool.true for FIFO pool, false for LIFO pool.Note that fifo has higher overhead compared to lifo. |
| pool_size | all | Maximum number of socket connections. Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.|
| min_idle_conns | all | Minimum number of idle connections which is useful when establishing new connection is slow. |
| max_conn_age | all | Connection age at which client retires (closes) the connection.Default is to not close aged connections. |
| pool_timeout | all | Amount of time client waits for connection if all connections are busy before returning an error.Default is ReadTimeout + 1 second. |
| idle_timeout | all | Amount of time after which client closes idle connections.Should be less than server's timeout.Default is 5 minutes. -1 disables idle timeout check. |
| idle_check_frequency | all | Frequency of idle checks made by idle connections reaper.Default is 1 minute. -1 disables idle connections reaper,but idle connections are still discarded by the client if IdleTimeout is set. |
| max_redirects | 集群 | The maximum number of retries before giving up. Command is retried on network errors and MOVED/ASK redirects. |
| read_only | 集群 | Enables read-only commands on slave nodes. |
| route_by_latency | 集群 | Allows routing read-only commands to the closest master or slave node.It automatically enables ReadOnly. |
| route_randomly | 集群 | Allows routing read-only commands to the random master or slave node.It automatically enables ReadOnly. |
| master_name | 哨兵 | The master name. |
| sentinel_password | 哨兵 | Sentinel password from "requirepass <password>" (if enabled) in Sentinel configuration, or, if SentinelUsername is also supplied, used for ACL-based authentication. |


## 数据库
数据库使用的是`gorm.io/gorm`库，GORM 官方支持的数据库类型有： MySQL, PostgreSQL, SQlite, SQL Server。  
同时，GORM 使用 `database/sql` 维护连接池

#### 使用：
```
import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/zhaoyang1214/ginco/framework/contract"
    "github.com/zhaoyang1214/ginco/framework/database"
)

// 表名为users
type User struct {
    ID           uint
    Name         string
}


// 使用默认连接，database.yaml文件中配置默认连接名（database.default）
defaultDb := a.GetI("db").(*database.Database)
user := User{}
defaultDb.First(&user)
fmt.Printf("%+v\n", user)

// 使用指定连接 - mysql2
dbCon := a.GetI("db").(contract.Database).Connection("mysql2")
user2 := User{}
dbCon.First(&user2)
fmt.Printf("%+v\n", user2)

// 使用指定连接 - sqlite
sqliteDb := a.GetI("db").(contract.Database).Connection("sqlite")
sqliteDb.Exec("DROP TABLE IF EXISTS `users`;CREATE TABLE `users` (`id` INTEGER, name TEXT, PRIMARY KEY(id));")
sqliteDb.Create(&User{
    Name: "test2",
})
user3 := User{}
sqliteDb.First(&user3)
fmt.Printf("%+v\n", user3)
```

#### 配置：
| 配置项 | 说明 |
| --- | --- |
| driver | 数据库驱动，支持mysql、sqlite、sqlserver、postgres |
| dsn | 数据库连接地址，当为单机时候配置 |
| write.*.dsn | 数据库写连接地址，当需要读写分离的时候配置，支持多读多写 |
| read.*.dsn | 数据库读连接地址，当需要读写分离的时候配置，支持多读多写 |
| conn_max_idle_time | 连接空闲的最大时间，单位小时 |
| conn_max_lifetime | 连接可复用的最大时间，单位小时 |
| max_idle_conns | 空闲连接池中连接的最大数量 |
| max_open_conns | 打开数据库连接的最大数量 |


 ## 缓存
缓存目前支持`redis`、`database`、`memory`三种驱动，其中`redis`、`database`分别基于ginco框架的`redis`和`database`服务，`memory`基于`bigcache`包。   

#### 使用
```
c := a.GetI("cache").(contract.Cache)
ctx := context.Background()
			
// 设置缓存
_ = c.Set(ctx, "test2", []byte("2222"), time.Minute)

// 获取缓存
v, err := c.Get(ctx, "test2")
fmt.Println(err)
if err != nil {
    fmt.Println(string(v))
}

// 删除缓存
_ = c.Delete(ctx, "test1", "test2")

// 判断缓存是否存在
c.Has(ctx, "test")

// 根据前缀清除缓存
_ = c.ClearPrefix(ctx, "te")

// 清空缓存
_ = c.Clear(ctx)
```

> 备注： database驱动需要创建缓存表，表名可配，表中需有key(string)、value(string)、expiration(time.Time)三个字段。

## 数据验证
有时候我们需要对数据进行验证，数据不一定来自表单等。这时可以使用独立的数据验证组件来验证数据。
#### 使用
```
type User struct {
	Name string `validate:"required"`
	Age int `validate:"gt=0"`
}

validate := a.GetI("validate").(*validator.Validate)
err := validate.Struct(&User{
    "",
    0,
})
if err != nil {
    for _, err := range err.(validator.ValidationErrors) {
        fmt.Println(err)
    }
}
```

> 表单验证可以使用 gin.Context 的 ShouldBind 等系列方法

待更。。。

