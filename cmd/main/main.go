package main

import (
	"fmt"

	"github.com/dsrosen6/addigy-command/internal/addigy"
)

func main() {
	if err := addigy.CheckRoot(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
