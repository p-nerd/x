package cmd

import (
	"os"

	"github.com/p-nerd/x/pkg/wos"
)

func Commit() {
	wos.ExecuteWithExitError("git", "add", "-A")
	msg := "update"
	if len(os.Args) >= 3 && os.Args[2] != "" {
		msg = os.Args[2]
	}
	wos.ExecuteWithExitError("git", "commit", "-m", msg)
}
