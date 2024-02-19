package cmd

import (
	"github.com/p-nerd/x/pkg/wos"
)

func Push() {
	Commit()
	wos.ExecuteWithExitError("git", "push")
}
