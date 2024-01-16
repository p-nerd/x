package cmd

import (
	"github.com/p-nerd/x/pkg/util"
)

func Up() {
	util.ExecuteWithExitError("docker", "compose", "up")
}
