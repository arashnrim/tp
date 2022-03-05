package cmd

import (
	"os"

	"github.com/arashnrim/tp/internal"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCommand)
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists all the locations configured.",
	Long: `When running the command, tp will look through the
config folder ($HOME/.tp) and locate set up locations. It will
then display these locations by their names and locations.

This command requires no arguments. To use the command, run ` +
		"`tp list`.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			color.New(color.FgRed).Add(color.Bold).Printf("Error: 0 arguments expected, received %d\n", len(args))
			os.Exit(1)
		} else {
			if err := internal.ListLocations(); err != nil {
				color.New(color.FgRed).Add(color.Bold).Printf("Error: %s\n", err)
				os.Exit(1)
			}
		}
	},
}
