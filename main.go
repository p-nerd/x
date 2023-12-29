package main

import (
	"os"

	"github.com/p-nerd/x/commands"
)

func main() {
	if os.Args[1] == "set" {
		commands.Set()
		return
	}
	commands.Root()
}
