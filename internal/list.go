package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func ListLocations() error {
	if err := ValidateConfigFolder(); err != nil {
		return err
	}

	files, err := os.ReadDir(filepath.Join(os.Getenv("HOME"), ".tp"))
	if err != nil {
		return err
	}

	listItems := [][]string{}
	lengthiestName := 0
	lengthiestPath := 0
	for _, file := range files {
		// Ignores all files but .yaml ones
		if filepath.Ext(file.Name()) == ".yaml" {
			var parsedData ConfigFile

			data, rfErr := os.ReadFile(filepath.Join(os.Getenv("HOME"), ".tp", file.Name()))
			if rfErr != nil {
				return err
			}

			umErr := yaml.Unmarshal(data, &parsedData)
			if umErr != nil {
				return err
			}

			if len(strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))) > lengthiestName {
				lengthiestName = len(strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())))
			}

			if len(parsedData.Location) > lengthiestPath {
				lengthiestPath = len(parsedData.Location)
			}

			listItem := []string{strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())), parsedData.Location}
			listItems = append(listItems, listItem)
		}
	}

	// Pretty prints the results
	if len(listItems) > 1 {
		fmt.Println("Here's a list of all the locations set up:")
		fmt.Print("Name")
		fmt.Print(strings.Repeat(" ", (lengthiestName - len("Name"))))
		fmt.Print(" | ")
		fmt.Print("Location")
		fmt.Println(strings.Repeat(" ", (lengthiestPath - len("Location"))))
		fmt.Println(strings.Repeat("-", (lengthiestName-len("Name"))+(lengthiestPath-len("Location"))))
		for _, listItem := range listItems {
			name, location := listItem[0], listItem[1]
			fmt.Print(name)
			fmt.Print(strings.Repeat(" ", (lengthiestName - len(name) + 2)))
			fmt.Println(location)
		}
	} else {
		fmt.Println("There are no locations set up.")
		fmt.Println("Try adding a location with `tp add <name> <location>` and try again!")
	}

	return nil
}
