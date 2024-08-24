package config

import (
	"github.com/caarlos0/env/v10"
)

type dbConfig struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Dbname   string `env:"DB_DBNAME"`
	Port     string `env:"DB_PORT"`
	Zone     string `env:"DB_ZONE"`
}

func GetDBConfig() dbConfig {
	cfg := dbConfig{}
	env.Parse(&cfg)
	return cfg
}
