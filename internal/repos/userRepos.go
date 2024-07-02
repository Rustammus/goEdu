package repos

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/dto"
	"goEdu/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, dto *dto.CreateUserDTO) (pgtype.UUID, error)
	FindAll(ctx context.Context) ([]dto.ReadUserDTO, error)
	FindByUUID(ctx context.Context, uuid pgtype.UUID) (*dto.ReadUserDTO, error)
	Update(ctx context.Context, user *dto.UpdateUserDTO, uuid pgtype.UUID) (*dto.ReadUserDTO, error)
	Delete(ctx context.Context, uuid pgtype.UUID) error
	FindByEmailWithHash(ctx context.Context, email string) (*models.User, error)
}
