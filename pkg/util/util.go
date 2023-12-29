package util

import (
	"fmt"
	"os"
	"os/user"
)

func DevLog(message ...any) {
	// fmt.Print("[dev] ")
	// fmt.Println(message...)
}

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
