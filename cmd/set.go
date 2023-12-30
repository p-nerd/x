package cmd

import (
	"os"

	"github.com/p-nerd/x/conf"
	"github.com/p-nerd/x/pkg/xrc"
)

func Set() {
	name := conf.XRC_SCRIPT_NAME
	value := os.Args[2]

	err := xrc.Set(name, value)
	if err != nil {
		panic(err)
	}
}
