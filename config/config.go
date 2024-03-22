package config

import (
	"encoding/json"
	"os"
)

type ConfigElem struct {
	Size int    `json:"size"`
	Name string `json:"name"`
}

type Config struct {
	CommitWithSign bool `json:"commit_with_sign"`
	Types          []struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
		Icon string `json:"icon"`
	} `json:"types"`
	Body    *ConfigElem `json:"body"`
	Scope   *ConfigElem `json:"scope"`
	Summary *ConfigElem `json:"summary"`
	Branch  *ConfigElem `json:"branch"`
}

func NewConfigWith(file string) *Config {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}

	var c Config
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil
	}
	return &c
}
