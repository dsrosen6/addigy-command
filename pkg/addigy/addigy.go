package addigy

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/huh/spinner"
)

const (
	addigyFolder = "/Library/Addigy/"
	goAgent      = addigyFolder + "go-agent"
	statusPath   = addigyFolder + "ansible/status.json"
)

type command struct {
	mainCommand string
	args        []string
}

var policierRunCommand = command{
	mainCommand: goAgent,
	args:        []string{"policier", "run"},
}

// PolicierRunVerbose runs the Addigy policier with full verbose output.
func PolicierRunVerbose() error {
	c := policierRunCommand
	return commandWithOutput(c)
}

// PolicierRunWithSpinner runs the Addigy policier with a spinner. It runs the same command as PolicierRunVerbose, but with a spinner instead of the full output stream.
// Good for when you're running it in front of a user and want it to look less scary :)
func PolicierRunWithSpinner() error {
	c := policierRunCommand

	action := func() {
		commandWithoutOutput(c)
	}

	spinner.New().Type(spinner.Line).Title(" Running Addigy Policy...this may take a few minutes, or more.").Action(action).Run()

	fmt.Println("Addigy Policy run complete.")
	return nil
}

// PolicierRun starts the Addigy policier, and just lets you know if it started.
func PolicierRun() error {
	c := command{
		mainCommand: "launchctl",
		args:        []string{"start", "com.addigy.policier"},
	}
	if err := commandWithOutput(c); err != nil {
		return errors.New("could not start policier - consider running in full verbose mode (sudo addigy run)")
	}

	fmt.Println("Policy run started.")
	return nil
}

// ResetPolicyProgress resets all policy progress by deleting the status.json file. Next time the policy runs, it'll create a new status.json file.
func ResetPolicyProgress() error {
	err := os.Remove(statusPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No policy progress to reset.")
			return nil
		}
		return fmt.Errorf("could not reset policy progress: %w", err)
	}
	fmt.Println("Successfully reset policy progress.")
	return nil
}

// PolicierInstall installs an item using the policier.
//
// For example, you can install a Smart Software items by going to the edit page for the item and grabbing the long string of characters
// after "edit/"
//
// Example usage: addigy.PolicierInstall(a7f8s0ab-0m13-2819-858t-c30de1e2uj13)
func PolicierInstall(item string) error {
	c := command{
		mainCommand: goAgent,
		args:        []string{"policier", "install", item},
	}

	return commandWithOutput(c)
}
