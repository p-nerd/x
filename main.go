package main

import (
	"os"

	"github.com/p-nerd/x/cmd"
)

var cmds = map[string]func(){
	"version": cmd.Version,
	"help":    cmd.Help,
	"--help":  cmd.Help,
	"-h":      cmd.Help,
	"set":     cmd.Set,
	"up":      cmd.Up,
	"git-all": cmd.GA,
	"ga":      cmd.GA,
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
