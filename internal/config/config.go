package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	ConfigPath     string `json:"-"`
	StandartEditor string `json:"standart_editor"`
}

func New(customConfigPath string) Config {
	dir := "lsgo"
	configFileName := "config.json"

	configBase, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting standard config path:", err)
		configBase = "./"
	}

	if customConfigPath != "" {
		fmt.Println("Set custom config path")
		configBase = customConfigPath
	}

	configDir := filepath.Join(configBase, dir)

	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		fmt.Println("Error creating config dir:", err)
	}

	configPath := filepath.Join(configDir, configFileName)

	_, existError := os.Stat(configPath)
	if existError == nil {
		fmt.Println("Config already exists")
	} else if os.IsNotExist(existError) {
		fmt.Println("Config not exists, creating...")

		_, err := os.Create(configPath)
		if err != nil {
			fmt.Println("Error creating config file:", err)
		}
	} else {
		fmt.Println("Error checking config:", existError)
	}

	return Config{
		ConfigPath:     configPath,
		StandartEditor: "hx",
	}
}

func (c *Config) LoadConfig() {
	path := c.ConfigPath

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error open config file", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading config file", err)
		return
	}

	if len(data) == 0 {
		c.StandartEditor = "hx"
		return
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.StandartEditor = cfg.StandartEditor
}
