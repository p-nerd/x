package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/p-nerd/x/conf"
	"github.com/p-nerd/x/pkg/util"
	"github.com/p-nerd/x/pkg/xrc"
)

func filesNameInDir(directoryPath string) ([]string, error) {
	var files []string
	if directoryPath == "" {
		return files, nil
	}
	util.DevLog("Reading directory: " + directoryPath)
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Printf("error reading directory %s: %v\n", directoryPath, err)
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func isScriptExist(files []string, name string) bool {
	for _, file := range files {
		if file == name {
			return true
		}
	}
	return false
}

func splitPath(path string) []string {
	dirs := strings.Split(path, "/")
	paths := []string{dirs[0]}
	for i, dir := range dirs {
		if i == 0 {
			continue
		}
		newPath := paths[len(paths)-1] + "/" + dir
		paths = append(paths, newPath)
	}
	return paths
}

func executeScript(scriptPath string, args ...string) error {
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Println("Script file does not exist:", scriptPath)
		return err
	}
	cmd := exec.Command(scriptPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing script:", err)
		return err
	}
	return nil
}

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getScriptNameAndArgs() (string, []string) {
	name, err := xrc.Get(conf.SCRIPT_NAME)
	if err != nil {
		name = "x.sh"
	}
	args := os.Args[1:]
	if args[0] == "-s" {
		name = args[1]
		args = args[2:]
	}
	return name, args
}

func Root() {
	scriptName, scriptArgs := getScriptNameAndArgs()
	util.DevLog("script name:", scriptName)
	util.DevLog("script args:", scriptArgs)

	currentDir, err := util.CurrentWorkingDirPath()
	util.DevLog("Current working directory is: " + currentDir)
	fatal(err)

	paths := splitPath(currentDir)
	util.DevLog("All splitted paths is: ", paths)

	for i := len(paths) - 1; i >= 0; i-- {
		path := paths[i]

		files, err := filesNameInDir(path)
		fatal(err)

		if isScriptExist(files, scriptName) {
			util.ChangeDir(path)
			executeScript(path+"/"+scriptName, scriptArgs...)
			util.ChangeDir(currentDir)
			return
		}
		util.DevLog(scriptName, "is not exit in", path)
	}
	fmt.Println(scriptName, "is not exit in any directory on the path")
}
