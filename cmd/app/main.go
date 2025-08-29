package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vasiliy-maslov/wallet-app/internal/config"
	"github.com/vasiliy-maslov/wallet-app/pkg/postgres"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("cannot init config", slog.Any("error", err))
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.Application.LogLevel}))
	slog.SetDefault(logger)

	dsnStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=false",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Name,
	)

	db, err := postgres.NewPostgresDB(context.Background(), dsnStr)
	if err != nil {
		slog.Error("cannot connect to db", slog.Any("error", err))
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Application started")
}
