package addigy

import (
	"errors"
	"fmt"
	"os/user"
)

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
