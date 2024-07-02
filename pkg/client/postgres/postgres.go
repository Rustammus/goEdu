package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"goEdu/internal/config"
	"goEdu/pkg/logging"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, cf config.Storage) (connect *pgx.Conn, err error) {
	logger := logging.GetLogger()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cf.Username, cf.Password, cf.Host, cf.Port, cf.Database)
	maxAttempts := 5

	for maxAttempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		connect, err = pgx.Connect(ctx, dsn)
		if err != nil {
			logger.Error("Failed to connect to database. Try again...")
			maxAttempts--
			time.Sleep(time.Second)
			continue
		}
		return connect, err
	}
	return nil, err
}

func NewPool(ctx context.Context, cf config.Storage) (pool *pgxpool.Pool, err error) {
	logger := logging.GetLogger()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cf.Username, cf.Password, cf.Host, cf.Port, cf.Database)
	maxAttempts := 5
	logger.Info("Connection to database")

	for maxAttempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			logger.Error("Failed to connect to database. Try again...")
			maxAttempts--
			time.Sleep(time.Second)
			continue
		}
		return pool, err
	}

	return pool, err
}
