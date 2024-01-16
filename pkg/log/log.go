package log

import (
	"fmt"

	"github.com/fatih/color"
)

func DevLog(message ...any) {
	// fmt.Print("[dev] ")
	// fmt.Println(message...)
}

var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()

func Green(message ...any) {
	fmt.Print(green(message...))
}

func Yellow(message ...any) {
	fmt.Print(yellow(message...))
}
