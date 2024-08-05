package main

import (
	"fmt"

	"github.com/dsrosen6/addigy-command/internal/cli"
	"github.com/dsrosen6/addigy-command/pkg/addigy"
)

func main() {
	if err := addigy.CheckRoot(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := cli.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
