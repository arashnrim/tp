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
			fmt.Println()
			color.New(color.FgRed).Add(color.BgWhite).Add(color.Bold).Printf("❌ %s failed to run.", command)
			fmt.Println("\n")
			return err
		}

		if err = cmd.Start(); err != nil {
			fmt.Println()
			color.New(color.FgRed).Add(color.BgWhite).Add(color.Bold).Printf("❌ %s failed to run.", command)
			fmt.Println("\n")
			return err
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			output := scanner.Text()
			fmt.Println(output)
		}
		cmd.Wait()
		fmt.Println()
		color.New(color.FgGreen).Add(color.BgWhite).Add(color.Bold).Printf("✓ %s ran successfully.", command)
		fmt.Println("\n")
	}

	return nil
}