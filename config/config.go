package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	databaseUrl := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse DATABASE_URL: %v\n", err)
		os.Exit(1)
	}

	DB, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

func Close() {
	DB.Close()
}
