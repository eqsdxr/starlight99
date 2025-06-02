package main

import (
	"github.com/BurntSushi/toml"
	"os"
)

var configPath string = "/home/ivan/.local/share/starlight99.toml"

type Config struct {
	PlayerName string
	TotalScore int
}

func saveConfig(cfg Config, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(cfg)
}

func readConfig(path string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	return cfg, err
}
