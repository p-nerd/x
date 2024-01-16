package cmd

import (
	"github.com/p-nerd/x/pkg/wos"
)

func Up() {
	wos.ExecuteWithExitError("docker", "compose", "up")
}
