package cmd

import (
	"os"

	"github.com/arashnrim/tp/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(toCommand)
}

var toCommand = &cobra.Command{
	Use:   "to <name>",
	Short: "Teleports to a set-up location.",
	Long: `When running the command, tp will handle the changing
of directories and automatically run the tasks you have configured.
If no commands have been configured in the location's config file,

This command requires one argument: the name of the location you
wish to teleport to. To use the command, run ` + "`tp to <name>`.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			color.New(color.FgRed).Add(color.Bold).Printf("Error: 1 argument expected, received %d\n", len(args))
		} else {
			name := args[0]
			if err := internal.TeleportToLocation(name); err != nil {
				color.New(color.FgRed).Add(color.Bold).Printf("Error: %s\n", err)
				os.Exit(1)
			}
		}
	},
}
