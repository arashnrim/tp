package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

func TeleportToLocation(name string) error {
	if err := ValidateConfigFolder(); err != nil {
		return err
	}

	if err := VerifyNameExists(name); err == nil {
		return fmt.Errorf("location `%s` does not exist", name)
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

	if len(commands) == 0 {
		color.New(color.FgYellow).Add(color.Bold).Printf("Warning: no commands have been configured for %s\n", name)
		color.New(color.FgHiWhite).Printf("Have you configured %s with commands to run?", filepath.Join(os.Getenv("HOME"), ".tp", fmt.Sprintf("%s.yaml", name)))
	}

	for _, command := range commands {
		app := strings.Split(command, " ")[0]
		args := strings.Split(command, " ")[1:]
		cmd := exec.Command(app, args...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			color.New(color.FgRed).Add(color.Bold).Printf("\n❌ `%s` failed to run.\n\n", command)
			return err
		}

		if err = cmd.Start(); err != nil {
			color.New(color.FgRed).Add(color.Bold).Printf("\n❌ `%s` failed to run.\n\n", command)
			return err
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			output := scanner.Text()
			fmt.Println(output)
		}
		cmd.Wait()
		color.New(color.FgGreen).Add(color.Bold).Printf("\n✓ `%s` ran successfully.\n\n", command)
	}

	return nil
}
