package repo

import (
	"context"
	"fmt"

	"github.com/hsmtkk/sturdy-giggle/env"
	"github.com/jackc/pgx/v5"
)

func ConnectPostgres(ctx context.Context, pgConfig env.PostgresConfig) (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.Database)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Postgres Database: %w", err)
	}
	return conn, nil
}
