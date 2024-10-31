package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		body("x.sh", args)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func body(script string, args []string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	paths := split(currentDir)
	for i := len(paths) - 1; i >= 0; i-- {
		path := paths[i]
		files, err := files(path)
		if err != nil {
			log.Fatal(err)
		}
		if exist(files, script) {
			err := os.Chdir(path)
			if err != nil {
				log.Fatal(err)
			}
			scriptPath := path + "/" + script
			err = executablePermission(scriptPath)
			if err != nil {
				log.Fatal(err)
			}
			err = execute(scriptPath, args...)
			if err != nil {
				log.Fatal(err)
			}
			err = os.Chdir(currentDir)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	fmt.Println(script, "is not exit in any directory on the path")
}

func split(path string) []string {
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

func files(directoryPath string) ([]string, error) {
	var files []string
	if directoryPath == "" {
		return files, nil
	}
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func exist(files []string, name string) bool {
	for _, file := range files {
		if file == name {
			return true
		}
	}
	return false
}

func executablePermission(scriptPathOrCommand string) error {
	fileInfo, err := os.Stat(scriptPathOrCommand)
	if err != nil {
		return err
	}
	if fileInfo.Mode()&0111 == 0 {
		err = os.Chmod(scriptPathOrCommand, fileInfo.Mode()|0111)
		if err != nil {
			return err
		}
	}
	return nil
}

func execute(scriptPathOrCommand string, args ...string) error {
	fmt.Print("$ ")
	fmt.Println(scriptPathOrCommand, " ", strings.Join(args, " "))
	cmd := exec.Command(scriptPathOrCommand, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
