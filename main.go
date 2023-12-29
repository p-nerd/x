package main

import (
	"os"

	"github.com/p-nerd/x/cmd"
)

func main() {
	switch os.Args[1] {
	case "set":
		cmd.Set()
	default:
		cmd.Root()
	}
}
