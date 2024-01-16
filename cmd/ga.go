package cmd

import (
	"os"

	"github.com/p-nerd/x/pkg/util"
)

func GA() {
	util.ExecuteWithExitError("git", "add", "-A")
	msg := "update"
	if os.Args[2] != "" {
		msg = os.Args[2]
	}
	util.ExecuteWithExitError("git", "commit", "-m", msg)
	util.ExecuteWithExitError("git", "push")
}
