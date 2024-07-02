package repos

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"goEdu/internal/crud"
)

type Repositories struct {
	User UserRepository
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User: crud.NewUserCRUD(pool),
	}
}
