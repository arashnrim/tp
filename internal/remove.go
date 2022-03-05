package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func RemoveLocation(name string) error {
	if err := ValidateConfigFolder(); err != nil {
		return err
	}

	if err := VerifyNameExists(name); err == nil {
		return fmt.Errorf("location `%s` does not exist", name)
	}

	if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".tp", fmt.Sprintf("%s.yaml", name))); err != nil {
		return err
	}

	color.New(color.FgGreen).Add(color.Bold).Printf("The location `%s` has now been deleted.", name)

	return nil
}
