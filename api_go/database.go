package main

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDatabase() error {
	// Load .env when running from the project root or from the api_go directory.
	_ = godotenv.Load()
	_ = godotenv.Load("../.env")

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return errors.New("DATABASE_URL is not set")
	}

	var err error
	DB, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return err
	}

	err = DB.Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func CreateTables(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS metrics (
		id SERIAL PRIMARY KEY,
		cpu_percent DOUBLE PRECISION NOT NULL,
		memory_percent DOUBLE PRECISION NOT NULL,
		disk_percent DOUBLE PRECISION NOT NULL,
		timestamp DOUBLE PRECISION NOT NULL
	);
	`

	_, err := DB.Exec(ctx, query)
	return err
}

func SaveMetric(ctx context.Context, metric Metric) error {
	query := `
	INSERT INTO metrics (
		cpu_percent,
		memory_percent,
		disk_percent,
		timestamp
	)
	VALUES ($1, $2, $3, $4);
	`

	_, err := DB.Exec(
		ctx,
		query,
		metric.CPUPercent,
		metric.MemoryPercent,
		metric.DiskPercent,
		metric.Timestamp,
	)

	return err
}
