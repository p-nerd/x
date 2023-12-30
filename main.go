package main

import (
	"os"

	"github.com/p-nerd/x/cmd"
)

var cmds = map[string]func(){
	"version": cmd.Version,
	"help":    cmd.Help,
	"set":     cmd.Set,
}

func main() {
	if len(os.Args) == 1 {
		cmd.Root()
		return
	}
	if cmdFunc, ok := cmds[os.Args[1]]; ok {
		cmdFunc()
	} else {
		cmd.Root()
	}
}
