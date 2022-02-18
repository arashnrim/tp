package internal

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func ConvertToAbsolutePath(location string) string {
	if !path.IsAbs(location) {
		location = filepath.Join(os.Getenv("HOME"), location)
	}
	return location
}

func AddLocation(name string, location string) error {
	if err := VerifyLocationFolder(location); err != nil {
		return err
	}
	location = ConvertToAbsolutePath(location)

	if err := VerifyNameExists(name); err != nil {
		return err
	}

	// Compiling the payload to be stored in the config file
	payload := ConfigFile{Location: location, Commands: []string{}}

	// Storing the data into the config file
	data, err := yaml.Marshal(&payload)
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(os.Getenv("HOME"), ".tp", fmt.Sprintf("%s.yaml", name)), data, 0777); err != nil {
		return err
	}

	fmt.Printf("All done! tp has set up the location; to configure, go to ~/.tp/%s.yaml.\n", name)
	fmt.Printf("You may now use `tp to %s` to instantly teleport and run the commands.\n", name)

	return nil
}

func VerifyLocationFolder(location string) error {
	location = ConvertToAbsolutePath(location)
	if _, err := os.Stat(location); os.IsNotExist(err) {
		return err
	}
	return nil
}
