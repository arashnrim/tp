package cmd

import (
	"fmt"
	"log"

	"github.com/arashnrim/tp/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Adds one or more locations to teleport to",
	Long: `When running the command, tp will create a new config file
for this location at the config folder ($HOME/.tp). You may then edit
the contents of this file afterwards and add additional steps to run
from there.

A wizard will guide you through the creation of this file. To skip the
wizard, use -s (or --skip-wizard).`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal(fmt.Errorf("at least 1 location name expected, received 0"))
		} else if len(args) > 0 {
			for _, location := range args {
				if err := utils.VerifyLocationFolder(location); err != nil {
					log.Fatal(err)
				} else {
					if err := utils.StartAddLocationWizard(); err != nil {

					}
				}
			}
		}
	},
}
