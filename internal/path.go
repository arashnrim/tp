package internal

import (
	"errors"
	"os"
	"path/filepath"
)

func CheckHomeVariable() error {
	if _, exists := os.LookupEnv("HOME"); !exists {
		return errors.New("HOME environment variable is not defined")
	}
	return nil
}

func ValidateConfigFolder() error {
	folderPath := filepath.Join(os.Getenv("HOME"), ".tp")
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.Mkdir(folderPath, 0777); err != nil {
			return err
		}
	}
	return nil
}
