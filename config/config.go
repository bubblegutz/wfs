package config

import (
	"log"
	"strings"

	"os"
	"path/filepath"
)

var customPath string = ""

func SetConfigPath(path string) {
	customPath = path
}

func GetConfigPath() string {
	if customPath != "" {
		return customPath
	}
	if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
		return filepath.Join(xdgConfig, "wfs")
	}
	home := os.Getenv("HOME")
	if home == "" {
		log.Fatal("HOME environment variable not set")
	}
	return filepath.Join(home, ".config", "wfs")
}

func GetConfigFiles() ([]os.DirEntry, error) {
	rootPath := filepath.Join(GetConfigPath(), "root")
	files, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func LoadConfigs() (map[string][]byte, error) {
	rootPath := filepath.Join(GetConfigPath(), "root")
	files, err := os.ReadDir(rootPath)
	if err != nil {
		log.Fatal(err)
	}
	configs := make(map[string][]byte)
	for _, f := range files {
		file := filepath.Join(rootPath, f.Name())
		if _, err := os.Stat(file); !os.IsNotExist(err) && f.Name() != "lib" {
			data, err := os.ReadFile(file)
			if err != nil {
				log.Fatal(err)
			}
			configs[strings.Replace(f.Name(), ".js", "", -1)] = data
		}
	}
	return configs, nil
}
