package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoyang1214/ginco/app/http/controller"
	"github.com/zhaoyang1214/ginco/app/http/middleware"
	"github.com/zhaoyang1214/ginco/framework/contract"
)

func Register(app contract.Application) {
	routerServer, err := app.Get("router")
	if err != nil {
		panic(err)
	}
	r := routerServer.(*gin.Engine)

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", controller.Index)
	r.GET("/name", controller.Name(app))

	authMiddleware := middleware.JWT(app)
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/logout", authMiddleware.LogoutHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/name", controller.Name(app))
	}

	registerApi(r)
}
