package db

import (
	"careerpilot/logger"
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func Connect() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Warn("no .env file found, relying on system environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		logger.Log.Error("DATABASE_URL is not set")
		os.Exit(1)
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		logger.Log.Error("unable to create connection pool", "error", err)
		os.Exit(1)
	}

	if err := pool.Ping(context.Background()); err != nil {
		logger.Log.Error("unable to reach database", "error", err)
		os.Exit(1)
	}

	Pool = pool
	logger.Log.Info("connected to database successfully")
}
