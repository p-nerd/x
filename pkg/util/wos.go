package util

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
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

func Execute(scriptPathOrCommand string, args ...string) error {
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
