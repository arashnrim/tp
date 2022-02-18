package cmd

import (
	"fmt"
	"log"

	"github.com/arashnrim/tp/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add <name> <location>",
	Short: "Adds one or more locations to teleport to",
	Long: `When running the command, tp will create a new config file
for this location at the config folder ($HOME/.tp). You may then edit
the contents of this file afterwards and add additional steps to run
from there. If the location is invalid, an error will be thrown.

This command requires two arguments: a name and a location. To use the
command, run tp add <name> <location>.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal(fmt.Errorf("2 arguments expected, received %d", len(args)))
		} else {
			// TODO: Implement adding of location
			name := args[0]
			location := args[1]
			if err := internal.AddLocation(name, location); err != nil {
				log.Fatal(err)
			}
		}
	},
}
