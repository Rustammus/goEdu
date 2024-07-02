package crud

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"goEdu/internal/config"
	"goEdu/pkg/client/postgres"
	"goEdu/pkg/logging"
)

var ConnPool *pgxpool.Pool

func init() {
	logger := logging.GetLogger()

	conf := config.GetConfig()
	pool, err := postgres.NewPool(context.TODO(), conf.Storage)
	if err != nil {
		logger.Fatalf("!!! Creating new client ERROR !!! \n Error: %s", err.Error())
	}
	ConnPool = pool
}
