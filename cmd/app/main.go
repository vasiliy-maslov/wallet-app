package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/vasiliy-maslov/wallet-app/internal/app"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	dsnStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=false",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Name,
	)

	db, err := postgres.NewPostgresDB(ctx, dsnStr)
	if err != nil {
		slog.Error("cannot connect to db", slog.Any("error", err))
		os.Exit(1)
	}

	application := app.New(logger, db)

	if err := application.Run(ctx); err != nil {
		slog.Error("application run failed", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("application stopped gracefully")

}
