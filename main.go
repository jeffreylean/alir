package main

import (
	"fmt"
	"os"

	"github.com/jeffreylean/alir/cmd"
)

func main() {
	command := cmd.New()

	if err := command.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
