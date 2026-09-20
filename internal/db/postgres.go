package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Firuz43/order-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(cfg *config.Config) (*pgxpool.Pool, error) {
	//Construct the connection string DSN
	//dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=" + cfg.DBSSLMode
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	//Parse configuration string into pgxpool.Config
	poolConfig, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse database config: %w", err)
	}

	//Configure connection pool settings

	// Configure Connection Pool Settings
	poolConfig.MaxConns = 25                      // Maximum active + idle connections
	poolConfig.MinConns = 5                       // Minimum warm connections kept open
	poolConfig.MaxConnLifetime = 30 * time.Minute // Max age of a single connection
	poolConfig.MaxConnIdleTime = 15 * time.Minute // Close connections idle for 15+ minutes

	// Create context with a 5 second timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Intialize connection pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)

	if err != nil {
		return nil, fmt.Errorf("Unable to create a connection pool: %w", err)
	}

	// Ping the database to verify connection is alive
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("database ping failed %w", err)
	}

	log.Println("Successfully connected to PostgreSQL via pgxpool")
	return pool, nil
}
