# Ginco

## 介绍
Ginco是一个Golang框架，基于gin框架和cobra CLI库实现。 大部分服务基于契约，均可替换。

##### 目前实现的服务：

| 服务 | 别名 | 备注 |
| --- | --- | --- |
| config | - | 配置服务，基于contract.Config契约，底层使用viper.Viper |
| console | cmd | 命令行服务，使用cobra.Command |
| http | server、router | HTTP服务，底层使用gin.Engine |
| logger | log | 日志服务，基于contract.Logger契约，底层使用zap.Logger |

> 其他服务待实现，当然你也可以自己实现contract.Provider接口，然后注册服务即可


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
a.GetIgnore("config").(contract.Config).GetString("app.name")
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
config := a.GetIgnore("config").(contract.Config)
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
conf := a.GetIgnore("conf").(contract.Config)
```

## 服务提供者
实现`contract.Provider`接口，然后绑定到容器`a.Bind("serverName", server)`即可

## 契约
使用契约是为了解耦。目前约定了部分契约，后续还会增加

## 路由
路由使用的是`gin.Engine`，只要在`router.Register`中注册相应的路由即可。
```shell
router := a.GetIgnore("router").(*gin.Engine)
router.GET("/", func (c *gin.Context) {
	c.String(http.StatusOK, "Hello Ginco v"+a.Version()+"\n")
})
```

具体使用请参考https://github.com/gin-gonic/gin

## 中间件
在`router.Register`中注册相应的中间件即可。
```shell
router := a.GetIgnore("router").(*gin.Engine)
router.Use(gin.Logger(), gin.Recovery())
```

## 日志
日志服务，基于contract.Logger契约，底层使用zap.Logger。  

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
log := a.GetIgnore("log").(contract.Logger)
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

