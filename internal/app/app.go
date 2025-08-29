package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Log *slog.Logger
	DB  *pgxpool.Pool
}

func New(log *slog.Logger, db *pgxpool.Pool) *App {
	return &App{
		Log: log,
		DB:  db,
	}
}

func (a *App) Run(ctx context.Context) error {
	a.Log.Info("Application started")
	<-ctx.Done()
	a.Log.Info("Stopping application")
	return a.Stop()
}

func (a *App) Stop() error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	a.DB.Close()
	a.Log.Info("Database connection closed")

	return nil
}
