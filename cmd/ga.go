package cmd

import (
	"os"

	"github.com/p-nerd/x/pkg/wos"
)

func GA() {
	wos.ExecuteWithExitError("git", "add", "-A")
	msg := "update"
	if os.Args[2] != "" {
		msg = os.Args[2]
	}
	wos.ExecuteWithExitError("git", "commit", "-m", msg)
	wos.ExecuteWithExitError("git", "push")
}
