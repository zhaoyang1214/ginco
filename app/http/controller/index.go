package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoyang1214/ginco/framework/contract"
	"github.com/zhaoyang1214/ginco/framework/foundation/app"
	"net/http"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello Ginco v"+app.Get().Version()+"\n")
}

func Name(app contract.Application) func(*gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "My name is "+app.GetIgnore("config").(contract.Config).GetString("app.name")+"\n")
	}
}
