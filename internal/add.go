package internal

import (
	"os"
	"path"
	"path/filepath"
)

func AddLocation(location string) error {
	if err := VerifyLocationFolder(location); err != nil {
		return err
	}
	return nil
}

func VerifyLocationFolder(location string) error {
	if !path.IsAbs(location) {
		location = filepath.Join(os.Getenv("HOME"), location)
	}
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return err
	}
	return nil
}
