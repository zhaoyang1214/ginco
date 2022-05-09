package service

import "github.com/zhaoyang1214/ginco/framework/contract"

type Index struct {
	*Service
}

func NewIndex(app contract.Application) *Index {
	return &Index{
		&Service{
			app: app,
		},
	}
}

func (i Index) Name() string {
	return i.app.GetI("config").(contract.Config).GetString("app.name")
}
