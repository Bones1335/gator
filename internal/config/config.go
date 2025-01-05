package config

import (
	"bytes"
	"encoding/json"
	"os"
)

type Config struct {
	DbURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	readJSONConfig, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	decoder := json.NewDecoder(bytes.NewReader(readJSONConfig))
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homePath + configFileName, nil

}

func (c *Config) SetUser(user string) error {
	c.CurrentUser = user
	return write(*c)
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}

	writeErr := os.WriteFile(configFilePath, jsonData, 0666)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
