package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

const FILENAME = "configs/config.json"

func GetConfig() *Config {
	configFile, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		msg := fmt.Sprintf("Failed to open config file: %s", err)
		panic(msg)
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		msg := fmt.Sprintf("Failed to read config file: %s", err)
		panic(msg)
	}

	return &config
}
