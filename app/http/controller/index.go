package controller

import (
	"ginco/app/model"
	"ginco/app/service"
	"ginco/framework/contract"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {
	*Controller
	service *service.Index
}

func NewIndex(app contract.Application) *Index {
	return &Index{
		Controller: &Controller{
			app: app,
		},
		service: service.NewIndex(app),
	}
}

func (i Index) Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello Ginco v"+i.app.Version()+"\n")
}

func (i Index) Name(c *gin.Context) {
	if userVal, ok := c.Get("user"); ok {
		user := userVal.(*model.User)
		c.String(http.StatusOK, "Hello userid %d username %s ", user.ID, user.Name)
		return
	}

	c.String(http.StatusOK, "My name is "+i.service.Name()+"\n")
}

// Hello World
// @Summary test
// @Schemes
// @Description test
// @Tags
// @Accept json
// @Produce json
// @Success 200 {object} entity.ResultJSON{code=int,message=string,data=string} "helloworld"
// @Router /helloworld [get]
func (i Index) Helloworld(c *gin.Context) {
	i.responseJson(c, 0, nil, "Hello World")
	return
}
