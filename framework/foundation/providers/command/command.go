package command

import (
	"ginco/framework/contract"
	"github.com/spf13/cobra"
)

type Command struct {
}

var _ contract.Provider = (*Command)(nil)

func (c *Command) Build(container contract.Container, params ...interface{}) (interface{}, error) {
	rootCmd := &cobra.Command{}
	return rootCmd, nil
}
