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
	Name     string `env:"POSTGRES_DB" env-default:"waller_db"`
	User     string `env:"POSTGRES_USER" env-default:"user"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
}

type Config struct {
	Application ApplicationConfig
	Postgres    PostgresConfig
}

func New() (*Config, error) {
	err := godotenv.Load("/configs/config.env")
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
