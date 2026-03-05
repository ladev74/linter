package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ladev74/linter/internal/analyzer"
)

type Config struct {
	Analyzer analyzer.Config `yaml:"analyzer"`
}

func New(path string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return &cfg, nil
}
