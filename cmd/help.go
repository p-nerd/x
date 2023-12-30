package cmd

import "fmt"

func Help() {
	help := `Execute any script from anywhere on the path

COMMANDS
  x <...args>                   run the default script (default script is x.sh)
  x -s <script name> <...args>  run the specified script
  x set <script name>           change default script name`
	fmt.Println(help)
}
