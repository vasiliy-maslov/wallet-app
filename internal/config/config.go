package config

import (
	"fmt"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ApplicationConfig struct {
	Addr     string     `env:"ADDR" env-default:":8080"`
	LogLevel slog.Level `env:"LOG_LEVEL" env-default:"DEBUG"`
}

type PostgresConfig struct {
	Name     string `env:"DB" env-default:"waller_db"`
	User     string `env:"USER" env-default:"user"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST" env-default:"localhost"`
	Port     string `env:"PORT" env-default:"5432"`
}

type Config struct {
	Application ApplicationConfig
	Postgres    PostgresConfig `env-prefix:"POSTGRES_"`
}

func New() (*Config, error) {
	err := godotenv.Load("configs/config.env")
	if err != nil {
		return nil, fmt.Errorf("error: cannot load config: %w", err)
	}

	var cfg Config

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error: cannot read config: %w", err)
	}

	return &cfg, err
}
