package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Numbone/user-api/config"
	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	log.Println("Connected to the database")
	return conn, nil
}
