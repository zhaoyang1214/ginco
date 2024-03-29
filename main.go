package main

import (
	"fmt"
	"ginco/bootstrap"
	"ginco/framework/foundation/app"
	"runtime/debug"
)

func main() {
	// todo 异常捕获优化
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	a := bootstrap.InitApp()
	app.Set(a)
	if err := bootstrap.Run(a); err != nil {
		panic(err)
	}
}
