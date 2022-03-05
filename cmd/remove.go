package cmd

import (
	"os"

	"github.com/arashnrim/tp/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCommand)
}

var removeCommand = &cobra.Command{
	Use:   "remove <name>",
	Short: "Removes a location to teleport to",
	Long: `When running the command, tp will delete the config file
for this location at the config folder ($HOME/.tp). If the given
name is invalid, an error will be thrown.

This command requires one argument: the name of the location you
wish to delete. To use the command, run tp add <name>.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			color.New(color.FgRed).Add(color.Bold).Printf("Error: 1 argument expected, received %d\n", len(args))
			os.Exit(1)
		} else {
			name := args[0]
			if err := internal.RemoveLocation(name); err != nil {
				color.New(color.FgRed).Add(color.Bold).Printf("Error: %s\n", err)
				os.Exit(1)
			}
		}
	},
}
