package command

import (
	"github.com/spf13/cobra"
	"github.com/zhaoyang1214/ginco/framework/contract"
)

type Command struct {
}

func (c *Command) Build(container contract.Container, params ...interface{}) (interface{}, error) {
	rootCmd := &cobra.Command{}
	return rootCmd, nil
}
