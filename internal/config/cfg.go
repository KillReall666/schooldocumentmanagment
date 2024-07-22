package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type Config struct {
	Address string `env:"RUN_ADDRESS"`
	DBPath  string `env:"DATABASE_URL"`
}

const (
	defaultServer = "localhost:8080"
	defaultDBPath = "host=localhost port=5432 user=Mr8 password=Rammstein12! dbname=school_db sslmode=disable"
)

func New() (*Config, error) {
	cfg := Config{}

	flag.StringVar(&cfg.Address, "a", defaultServer, "server address [host:port]")
	flag.StringVar(&cfg.DBPath, "d", defaultDBPath, "db address string [host= port= user= password= dbname= sslmode= ")

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
