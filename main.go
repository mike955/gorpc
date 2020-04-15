package main

import (
	"gorpc/cmd"
	"os"
)

func main() {
	command := cmd.NewCmdGorpc()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
