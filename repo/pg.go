package repo

import (
	"context"
	"fmt"

	"github.com/cenkalti/backoff/v4"
	"github.com/hsmtkk/sturdy-giggle/env"
	"github.com/jackc/pgx/v5"
)

func ConnectPostgres(ctx context.Context, pgConfig env.PostgresConfig) (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.Database)
	var conn *pgx.Conn
	operation := func() error {
		var err error
		conn, err = pgx.Connect(ctx, url)
		return err
	}
	if err := backoff.Retry(operation, backoff.NewExponentialBackOff()); err != nil {
		return nil, fmt.Errorf("failed to connect Postgres Database: %w", err)
	}
	return conn, nil
}
