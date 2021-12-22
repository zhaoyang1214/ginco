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
使用 `go run main.go help start`可以看到其他参数
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

## 服务容器

## 服务提供者

## 契约

## 路由

## 中间件

## 日志

## 命令行

