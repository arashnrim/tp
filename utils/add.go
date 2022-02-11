package utils

import (
	"os"
	"path/filepath"
)

func ValidateConfigFolder() error {
	folderPath := filepath.Join(os.Getenv("HOME"), ".tp")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, 0777); err != nil {
			return err
		}
	}
	return nil
}
