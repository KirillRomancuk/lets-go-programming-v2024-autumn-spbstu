package config

import (
	"errors"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

var (
	errParseConfig = errors.New("parse config failed")
)

type Config struct {
	// Server config
	// Storage config
}

func Unmarshal(r io.Reader) (Config, error) {
	var cfg Config

	if err := yaml.NewDecoder(r).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("%w: %w", errParseConfig, err)
	}

	return cfg, nil
}
