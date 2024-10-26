package wos

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/p-nerd/x/pkg/log"
)

func ChangeDir(newDir string) error {
	err := os.Chdir(newDir)
	return err
}

func CurrentWorkingDirPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return "", err
	}
	return currentDir, nil
}

func ReadFile(filePath string) (string, error) {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(fileContents), nil
}

func WriteFile(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func IsFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func UserHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.HomeDir, nil
}

func ExecutablePermission(scriptPathOrCommand string) error {
	// Get file info
	fileInfo, err := os.Stat(scriptPathOrCommand)
	if err != nil {
		return err
	}

	// Check if file has executable permission
	if fileInfo.Mode()&0111 == 0 {
		// Add executable permission
		err = os.Chmod(scriptPathOrCommand, fileInfo.Mode()|0111)
		if err != nil {
			return err
		}
	}
	return nil
}

func Execute(scriptPathOrCommand string, args ...string) error {
	log.Green("$ ")
	log.Yellow(scriptPathOrCommand, " ", strings.Join(args, " "), "\n")
	cmd := exec.Command(scriptPathOrCommand, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing:", err)
		return err
	}
	return nil
}

func ExecuteWithExitError(scriptPathOrCommand string, args ...string) {
	err := Execute(scriptPathOrCommand, args...)
	if err != nil {
		os.Exit(1)
	}
}
