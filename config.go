package main

import (
	"os"
	"github.com/BurntSushi/toml"
)

type Config struct {
	PlayerName, Difficulty string
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
