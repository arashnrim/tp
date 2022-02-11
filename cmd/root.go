package cmd

import (
	"os"

	"github.com/arashnrim/tp/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tp",
	Short: " ⚡️ Teleport to your projects and run tasks in an instant.",
	Long: `tp is a simple tool that allows you to quickly get started
with your work by handling the change of directory and running the
tasks you set it up to do automatically.

To begin, use tp add <location>. This must be a valid location on your
computer; otherwise, the program will return an error. To remove an
existing location, use tp remove <location>. For more information, type
tp help.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	utils.ValidateConfigFolder()
}
