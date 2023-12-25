package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getCurrentWorkingDirPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return "", err
	}
	return currentDir, nil
}

func filesNameInDir(directoryPath string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %v", directoryPath, err)
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

func printSlice(s []string) {
	for _, val := range s {
		fmt.Println(val)
	}
}

func executeScript(scriptPath string, args ...string) {
	// Check if the file exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Println("Script file does not exist:", scriptPath)
		return
	}

	// Create a new exec.Cmd instance with the script path and arguments
	cmd := exec.Command(scriptPath, args...)

	// Set the command's standard output and error to os.Stdout and os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing script:", err)
		return
	}

	fmt.Println("Script executed successfully.")
}

func changeDir(newDir string) error {
	err := os.Chdir(newDir)
	return err
}

func main() {
	scriptName := "x.sh"

	currentDir, _ := getCurrentWorkingDirPath()
	paths := splitPath(currentDir)

	for i := len(paths) - 1; i >= 0; i-- {
		path := paths[i]

		files, _ := filesNameInDir(path)
		isExist := isScriptExist(files, scriptName)

		if isExist {
			changeDir(path)
			scriptArgs := os.Args[1:]
			executeScript(path+"/"+scriptName, scriptArgs...)
			changeDir(currentDir)
			return
		}
	}
}
