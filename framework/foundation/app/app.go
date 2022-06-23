package app

import (
	"ginco/framework/contract"
)

var app contract.Application

func Set(a contract.Application) {
	app = a
}

func Get() contract.Application {
	return app
}
