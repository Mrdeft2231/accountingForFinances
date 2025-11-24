package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Postgres Postgres `yaml:"postgres"`
	}
	Postgres struct {
		User     string `yaml:"pg_user"`
		Password string `yaml:"pg_password"`
		Database string `yaml:"pg_db"`
		URL      string `yaml:"pg_url"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("../config/config.yaml", cfg); err != nil {
		return nil, err
	}
	fmt.Println("Сработал", cfg)
	return cfg, nil
}
