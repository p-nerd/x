package cmd

import "github.com/p-nerd/x/pkg/util"

func Up() {
	util.Execute("docker", "compose", "up")
}
