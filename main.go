package main

import (
	"fmt"
	"gorpc/cmd"
	"os"
)

func main()  {
	//command := cmd.NewGorpcCommand()
	//if err := command.Execute(); err != nil {
	//	os.Exit(1)
	//}
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}