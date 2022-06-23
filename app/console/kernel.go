package console

import (
	"ginco/app/console/command/version"
	"ginco/framework/contract"
	"github.com/spf13/cobra"
)

func Register(a contract.Application) {
	cmdServer, err := a.Get("cmd")
	if err != nil {
		panic(err)
	}

	cmd := cmdServer.(*cobra.Command)

	cmd.AddCommand(version.Command(a))
}
