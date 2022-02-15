package internal

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func ConvertToAbsolutePath(location string) string {
	if !path.IsAbs(location) {
		location = filepath.Join(os.Getenv("HOME"), location)
	}
	return location
}

func AddLocation(location string) error {
	if err := VerifyLocationFolder(location); err != nil {
		return err
	}
	location = ConvertToAbsolutePath(location)

	// Compiling the payload to be stored in the config file
	payload := ConfigFile{Location: location, Commands: []string{}}

	// Converting / to - for the file names
	trimmedConfigFileName := strings.Trim(location, "/")
	splitConfigFileName := strings.Split(trimmedConfigFileName, "/")
	configFileName := splitConfigFileName[len(splitConfigFileName)-1]

	// Storing the data into the config file
	data, err := yaml.Marshal(&payload)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(os.Getenv("HOME"), ".tp", fmt.Sprintf("%s.yaml", configFileName)), data, 0777); err != nil {
		return err
	}

	fmt.Printf("All done! tp has set up the location; to configure, go to ~/.tp/%s.yaml.\n", configFileName)
	fmt.Printf("You may now use tp to %s to instantly teleport and run the commands.\n", location)

	return nil
}

func VerifyLocationFolder(location string) error {
	location = ConvertToAbsolutePath(location)
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return err
	}
	return nil
}
