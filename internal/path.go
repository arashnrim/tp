package internal

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type ConfigFile struct {
	Location string   `yaml:location`
	Commands []string `yaml:commands`
}

func CheckHomeVariable() error {
	if _, exists := os.LookupEnv("HOME"); !exists {
		return errors.New("$HOME environment variable is not defined")
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

func VerifyNameExists(name string) error {
	if _, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".tp", fmt.Sprintf("%s.yaml", name))); !os.IsNotExist(err) {
		return fmt.Errorf("location of name `%s` already exists", name)
	}
	return nil
}
