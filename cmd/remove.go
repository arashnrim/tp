package cmd

import (
	"fmt"
	"log"

	"github.com/arashnrim/tp/internal"
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
			log.Fatal(fmt.Errorf("1 argument expected, received %d", len(args)))
		} else {
			name := args[0]
			if err := internal.RemoveLocation(name); err != nil {
				log.Fatal(err)
			}
		}
	},
}
