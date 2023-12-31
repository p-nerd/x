package cmd

import (
	"fmt"

	"github.com/p-nerd/x/conf"
)

func Version() {
	fmt.Println("x version", conf.VERSION)
}
