package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zhaoyang1214/ginco/app/http/controller"
	"github.com/zhaoyang1214/ginco/app/http/middleware"
	docs "github.com/zhaoyang1214/ginco/docs"
	"github.com/zhaoyang1214/ginco/framework/contract"
)

func Register(app contract.Application) {
	routerServer, err := app.Get("router")
	if err != nil {
		panic(err)
	}
	r := routerServer.(*gin.Engine)

	r.Use(gin.Logger(), gin.Recovery())

	i := controller.NewIndex(app)

	r.GET("/", i.Index)
	r.GET("/name", i.Name)

	authMiddleware := middleware.JWT(app)
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/logout", authMiddleware.LogoutHandler)
	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/name", i.Name)
	}

	docs.SwaggerInfo_swagger.BasePath = "/example"
	example := r.Group("/example")
	{
		example.GET("/helloworld", i.Helloworld)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	registerApi(r)
}
