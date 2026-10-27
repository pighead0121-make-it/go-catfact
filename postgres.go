package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	dbpool *pgxpool.Pool
}

func NewPostgresStore(ctx context.Context, connStr string) (*PostgresStore, error) {
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(ctx); err != nil {
		return nil, err
	}

	return &PostgresStore{
		dbpool: dbpool,
	}, nil
}

func (s *PostgresStore) Init(ctx context.Context) error {
	query := `
		CREATE TABLE IF NOT EXISTS request_logs (
			id SERIAL PRIMARY KEY,
			request_id INT NOT NULL,
			fact TEXT NOT NULL,
			err TEXT,
			created_at TIMESTAMPTZ NOT NULL
		);`

	_, err := s.dbpool.Exec(ctx, query)
	return err
}

func (s *PostgresStore) InsertRequestLog(ctx context.Context, reqID int, fact string, reqErr error, createdAt time.Time) {
	query := `
		INSERT INTO request_logs (request_id, fact, err, created_at)
		VALUES ($1, $2, $3, $4)
	`

	var errStr string
	if reqErr != nil {
		errStr = reqErr.Error()
	}

	_, err := s.dbpool.Exec(ctx, query, reqID, fact, errStr, createdAt)
	if err != nil {
		log.Printf("failed to write request log to db: %v", err)
	}
}
