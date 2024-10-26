package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/p-nerd/x/conf"
	"github.com/p-nerd/x/pkg/log"
	"github.com/p-nerd/x/pkg/wos"
	"github.com/p-nerd/x/pkg/xrc"
)

func filesNameInDir(directoryPath string) ([]string, error) {
	var files []string
	if directoryPath == "" {
		return files, nil
	}
	log.DevLog("Reading directory: " + directoryPath)
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

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getScriptNameAndArgs() (string, []string) {
	name, err := xrc.Get(conf.XRC_SCRIPT_NAME)
	if err != nil {
		name = conf.DEFAULT_SCRIPT
	}
	if len(os.Args) <= 1 {
		return name, []string{}
	}
	if os.Args[1] == "-s" {
		name = os.Args[2]
		args := []string{}
		if len(os.Args) >= 4 {
			args = os.Args[3:]
		}
		return name, args
	}
	return name, os.Args[1:]
}

func Root() {
	scriptName, scriptArgs := getScriptNameAndArgs()
	log.DevLog("script name:", scriptName)
	log.DevLog("script args:", scriptArgs)

	currentDir, err := wos.CurrentWorkingDirPath()
	log.DevLog("Current working directory is: " + currentDir)
	fatal(err)

	paths := splitPath(currentDir)
	log.DevLog("All splitted paths is: ", paths)

	for i := len(paths) - 1; i >= 0; i-- {
		path := paths[i]

		files, err := filesNameInDir(path)
		fatal(err)

		if isScriptExist(files, scriptName) {
			wos.ChangeDir(path)
			scriptNameWithPath := path + "/" + scriptName
			err := wos.ExecutablePermission(scriptNameWithPath)
			if err != nil {
				log.Yellow("Failed to give executable permission", err)
			}
			err = wos.Execute(scriptNameWithPath, scriptArgs...)
			wos.ChangeDir(currentDir)
			if err != nil {
				os.Exit(1)
			}
			return
		}
		log.DevLog(scriptName, "is not exit in", path)
	}
	fmt.Println(scriptName, "is not exit in any directory on the path")
}
