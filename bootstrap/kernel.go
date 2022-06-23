package bootstrap

import (
	"ginco/framework/contract"
	"ginco/framework/foundation/commands"
	"ginco/framework/foundation/providers/cache"
	"ginco/framework/foundation/providers/command"
	"ginco/framework/foundation/providers/config"
	"ginco/framework/foundation/providers/database"
	"ginco/framework/foundation/providers/elasticsearch"
	"ginco/framework/foundation/providers/http"
	"ginco/framework/foundation/providers/logger"
	"ginco/framework/foundation/providers/redis"
	"ginco/framework/foundation/providers/validate"
	"github.com/spf13/cobra"
)

func Run(a contract.Application) error {
	commandServer, err := a.Get("console")
	if err != nil {
		panic(err)
	}
	return commandServer.(*cobra.Command).Execute()
}

func registerBaseProviders(a contract.Application) {
	a.Bind("config", &config.Config{})
	a.Bind("console", &command.Command{})
	a.Bind("http", &http.Http{})
	a.Bind("logger", &logger.Logger{})
	a.Bind("redis", &redis.Redis{})
	a.Bind("database", &database.Database{})
	a.Bind("cache", &cache.Cache{})
	a.Bind("validate", &validate.Validate{})
	a.Bind("elasticsearch", &elasticsearch.Elasticsearch{})
}

func registerCoreAliases(a contract.Application) {
	aliases := map[string]string{
		"cmd":       "console",
		"server":    "http",
		"router":    "http",
		"log":       "logger",
		"db":        "database",
		"validator": "validate",
		"es":        "elasticsearch",
	}
	for k, v := range aliases {
		a.Alias(v, k)
	}
}

func registerConfigAliases(a contract.Application) {
	configServer, err := a.Get("config")
	if err != nil {
		panic(err)
	}
	conf := configServer.(contract.Config)

	aliases := conf.GetStringMapString("app.aliases")
	for k, v := range aliases {
		a.Alias(v, k)
	}
}

func registerBaseCommands(a contract.Application) {
	commandServer, err := a.Get("console")
	if err != nil {
		panic(err)
	}

	cmd := commandServer.(*cobra.Command)

	cmd.AddCommand(commands.HttpCommand(a))
	cmd.AddCommand(commands.HttpStopCommand(a))
	cmd.AddCommand(commands.HttpRestartCommand(a))
	cmd.AddCommand(commands.MigrateCommand(a))
	cmd.AddCommand(commands.MigrateRollbackCommand(a))
}
