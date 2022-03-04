package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

func TeleportToLocation(name string) error {
	if err := ValidateConfigFolder(); err != nil {
		return err
	}

	if err := VerifyNameExists(name); err == nil {
		return err
	}

	data, rfErr := os.ReadFile(filepath.Join(os.Getenv("HOME"), ".tp", fmt.Sprintf("%s.yaml", name)))
	if rfErr != nil {
		return rfErr
	}

	var parsedData ConfigFile
	if err := yaml.Unmarshal(data, &parsedData); err != nil {
		return err
	}

	location := parsedData.Location
	commands := parsedData.Commands
	if err := os.Chdir(location); err != nil {
		return err
	}

	for _, command := range commands {
		app := strings.Split(command, " ")[0]
		args := strings.Split(command, " ")[1:]
		cmd := exec.Command(app, args...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		if err = cmd.Start(); err != nil {
			return err
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			output := scanner.Text()
			fmt.Println(output)
		}
		cmd.Wait()
	}

	return nil
}
