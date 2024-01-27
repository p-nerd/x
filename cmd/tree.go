package cmd

import (
	"github.com/p-nerd/x/pkg/wos"
)

func Tree() {
	wos.ExecuteWithExitError("tree", "--gitignore")
}
