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

func VerifyLocationFolder(location string) error {
	var folderPath string
	if string([]rune(location)[0]) == "/" {
		folderPath = location
	} else {
		folderPath = filepath.Join(os.Getenv("HOME"), location)
	}
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return err
	}
	return nil
}

func StartAddLocationWizard(location string, folderPath string) error {
	// TODO: Implement wizard
	return nil
}
