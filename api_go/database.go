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

func GetLatestMetric(ctx context.Context) (Metric, error) {
	query := `
	SELECT
		id,
		cpu_percent,
		memory_percent,
		disk_percent,
		timestamp
	FROM metrics
	ORDER BY id DESC
	LIMIT 1;
	`

	var metric Metric

	err := DB.QueryRow(ctx, query).Scan(
		&metric.ID,
		&metric.CPUPercent,
		&metric.MemoryPercent,
		&metric.DiskPercent,
		&metric.Timestamp,
	)

	return metric, err
}

func GetMetricsHistory(ctx context.Context) ([]Metric, error) {
	query := `
	SELECT
		id,
		cpu_percent,
		memory_percent,
		disk_percent,
		timestamp
	FROM metrics
	ORDER BY id DESC
	LIMIT 25;
	`

	rows, err := DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []Metric

	for rows.Next() {
		var metric Metric

		err := rows.Scan(
			&metric.ID,
			&metric.CPUPercent,
			&metric.MemoryPercent,
			&metric.DiskPercent,
			&metric.Timestamp,
		)
		if err != nil {
			return nil, err
		}

		metrics = append(metrics, metric)
	}

	// Reverse the slice to match the Python API
	for i, j := 0, len(metrics)-1; i < j; i, j = i+1, j-1 {
		metrics[i], metrics[j] = metrics[j], metrics[i]
	}

	return metrics, rows.Err()
}
