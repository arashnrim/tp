package internal

import (
	"os"
	"path/filepath"
)

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
