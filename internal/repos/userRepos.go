package repos

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/dto"
	"goEdu/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, dto *dto.UserDTO) (pgtype.UUID, error)
	FindAll(ctx context.Context) ([]models.User, error)
	FindByUUID(ctx context.Context, uuid pgtype.UUID) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, uuid pgtype.UUID) error
}
