package xrc

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/p-nerd/x/pkg/wos"
)

type Config map[string]string

func parseXRC(content string) (Config, error) {
	config := make(Config)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}
		if strings.HasPrefix(line, "#") {
			continue // Skip comments lines
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config[key] = value
		}
	}

	return config, nil
}

func stringfy(config Config) string {
	s := ""
	for key, value := range config {
		s += fmt.Sprintf("%s=%s\n", key, value)
	}
	return s
}

func xrcFilepath() (string, error) {
	homeDir, err := wos.UserHomeDir()
	if err != nil {
		return "", err
	}
	xrcFilePath := filepath.Join(homeDir, ".xrc")
	return xrcFilePath, nil
}

func read() (Config, error) {
	xrcpath, err := xrcFilepath()
	if err != nil {
		return nil, err
	}
	content, err := wos.ReadFile(xrcpath)
	if err != nil {
		return nil, err
	}
	return parseXRC(content)
}

func write(config Config) error {
	xrcpath, err := xrcFilepath()
	if err != nil {
		return err
	}
	exist := wos.IsFileExist(xrcpath)
	if !exist {
		err = wos.CreateFile(xrcpath)
		if err != nil {
			return err
		}
	}
	s := stringfy(config)
	err = wos.WriteFile(xrcpath, s)
	return err
}

func Get(key string) (string, error) {
	config, err := read()
	if err != nil {
		return "", err
	}
	value := config[key]
	if value == "" {
		return "", errors.New("'" + key + "' key not found in .xrc file")
	}
	return value, nil
}

func Set(key string, value string) error {
	config, err := read()
	if err != nil {
		config = make(Config)
	}
	config[key] = value
	err = write(config)
	return err
}
