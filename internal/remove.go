package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemoveLocation(name string) error {
	if err := ValidateConfigFolder(); err != nil {
		return err
	}

	if err := VerifyNameExists(name); err == nil {
		return fmt.Errorf("location %s does not exist", name)
	}

	if err := os.Remove(filepath.Join(os.Getenv("HOME"), ".tp", fmt.Sprintf("%s.yaml", name))); err != nil {
		return err
	}

	fmt.Printf("The location %s has now been deleted.", name)

	return nil
}
