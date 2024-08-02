package addigy

import (
	"os"
	"os/exec"
)

func commandWithOutput(c command) error {
	cmd := exec.Command(c.mainCommand, c.args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func commandWithoutOutput(c command) error {
	cmd := exec.Command(c.mainCommand, c.args...)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
