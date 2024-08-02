package addigy

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

const (
	addigyFolder = "/Library/Addigy/"
	goAgent      = addigyFolder + "go-agent"
	statusPath   = addigyFolder + "ansible/status.json"
)

// runPolicier runs the policier binary with the given argument. Gives full output to terminal.
func policierCommand(args []string) error {
	a := []string{"policier"}
	a = append(a, args...)
	cmd := exec.Command(goAgent, a...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running policier: %w", err)
	}

	return nil
}

func CheckRoot() error {
	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	if u.Uid != "0" {
		return errors.New("command must be run as root - switch to root or add a sudo")
	}

	return nil
}
