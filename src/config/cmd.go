package config

import (
	"github.com/spf13/cobra"
	"github.com/wildanfaz/backendgolang2_week9/src/database/orm"
)

var initCommand = cobra.Command{
	Short: "backend golang week 9",
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
