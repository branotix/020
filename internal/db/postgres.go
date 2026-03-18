package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/branotix/p2p/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDatabase() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_NAME"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("❌ Unable to connect to database: %v\n", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("❌ Database ping failed: %v\n", err)
	}

	Pool = pool
	log.Println("✅ Connected to PostgreSQL with pgxpool")
}
