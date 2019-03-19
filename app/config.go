package app

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server     *ServerConfig
	EtcdServer *EtcdServerConfig
}

type ServerConfig struct {
	Address   string
	Port      int
	ClientTTL int64
}

type EtcdServerConfig struct {
	Urls []string
	Port int
	Path string
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
