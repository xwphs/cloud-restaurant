package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

// parse json config file
func ParseConfig(filePath string) (*Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var config Config
	r := bufio.NewReader(f)
	if err = json.NewDecoder(r).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
