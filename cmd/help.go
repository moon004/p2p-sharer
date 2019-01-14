package cmd

import (
	"github.com/moon004/p2p-sharer/tools"
	"github.com/spf13/cobra"
)

func helpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "help about any command",
		Long: `
help command will show the full details of the particular command

Example:
	` + tools.Args0() + ` help [command]`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}
