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

	jsonConfigPath := configFilePath + configFileName

	readJSONConfig, err := os.ReadFile(jsonConfigPath)
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

	return homePath, nil

}

