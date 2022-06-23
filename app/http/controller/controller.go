package controller

import (
	"ginco/app/entity"
	"ginco/framework/contract"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	app contract.Application
}

func (c Controller) responseJsonError(gc *gin.Context, msg string, code ...int) {
	reCode := http.StatusInternalServerError
	if len(code) > 0 {
		reCode = code[0]
	}
	gc.JSON(http.StatusOK, entity.ResultJSON{
		Code:    reCode,
		Message: msg,
	})
}

func (c Controller) responseJsonSuccess(gc *gin.Context, data interface{}, code ...int) {
	reCode := http.StatusOK
	if len(code) > 0 {
		reCode = code[0]
	}
	gc.JSON(http.StatusOK, entity.ResultJSON{
		Code:    reCode,
		Message: "OK",
		Data:    data,
	})
}

func (c Controller) responseJson(gc *gin.Context, code int, err error, data interface{}) {
	if err != nil {
		c.responseJsonError(gc, err.Error(), code)
		return
	}
	c.responseJsonSuccess(gc, data, code)
}
