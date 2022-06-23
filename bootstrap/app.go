package bootstrap

import (
	"ginco/app/console"
	"ginco/app/providers"
	"ginco/framework/contract"
	"ginco/framework/foundation"
	"os"
)

func InitApp() contract.Application {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var app contract.Application = foundation.NewApplication(path)

	registerBaseProviders(app)
	providers.Register(app)
	registerCoreAliases(app)
	registerConfigAliases(app)
	registerBaseCommands(app)
	console.Register(app)
	return app
}
