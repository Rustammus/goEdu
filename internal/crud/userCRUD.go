package crud

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/dto"
	"goEdu/internal/models"
	"goEdu/pkg/client/postgres"
)

type UserCRUD struct {
	client postgres.Client
}

func (u *UserCRUD) Create(ctx context.Context, createUserDTO *dto.CreateUserDTO) (pgtype.UUID, error) {
	q := `INSERT INTO public.users (name, phone, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING uuid`
	uuid := pgtype.UUID{}

	err := u.client.QueryRow(ctx, q,
		createUserDTO.Name, createUserDTO.Phone, createUserDTO.Email, createUserDTO.Password).Scan(&uuid)

	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func (u *UserCRUD) FindAll(ctx context.Context) ([]dto.ReadUserDTO, error) {

	q := `SELECT uuid, name, phone, email, crystals, is_blocked, updated_at, created_at FROM public.users`
	rows, err := u.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]dto.ReadUserDTO, 0)

	for rows.Next() {
		var user dto.ReadUserDTO
		err = rows.Scan(&user.UUID, &user.Name, &user.Phone, &user.Email, &user.Crystals, &user.IsBlocked, &user.UpdatedAt, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

func (u *UserCRUD) FindByUUID(ctx context.Context, uuid pgtype.UUID) (*dto.ReadUserDTO, error) {

	q := `SELECT name, phone, email, crystals, is_blocked, updated_at, created_at FROM public.users WHERE uuid=$1`

	user := dto.ReadUserDTO{UUID: uuid}

	err := u.client.QueryRow(ctx, q, uuid).Scan(&user.Name, &user.Phone, &user.Email, &user.Crystals, &user.IsBlocked, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		return &dto.ReadUserDTO{}, err
	}

	return &user, nil
}

func (u *UserCRUD) Update(ctx context.Context, user *dto.UpdateUserDTO, uuid pgtype.UUID) (*dto.ReadUserDTO, error) {

	q := `
		UPDATE public.users
		SET (name, phone, email) = ($2, $3, $4)
		WHERE uuid = $1 
		RETURNING uuid, name, phone, email, crystals, is_blocked, updated_at, created_at`

	readUser := dto.ReadUserDTO{}

	err := u.client.QueryRow(ctx, q, uuid, user.Name, user.Phone, user.Email).
		Scan(&readUser.UUID, &readUser.Name, &readUser.Phone, &readUser.Email,
			&readUser.Crystals, &readUser.IsBlocked, &readUser.UpdatedAt, &readUser.CreatedAt)

	if err != nil {
		return &dto.ReadUserDTO{}, err
	}

	return &readUser, nil
}

func (u *UserCRUD) Delete(ctx context.Context, uuid pgtype.UUID) error {
	//TODO implement me
	q := `DELETE FROM public.users WHERE uuid = $1 RETURNING uuid`
	err := u.client.QueryRow(ctx, q, uuid).Scan(&uuid)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserCRUD) FindByEmailWithHash(ctx context.Context, email string) (*models.User, error) {

	q := `SELECT uuid, name, phone, password_hash, crystals, is_blocked, updated_at, created_at FROM public.users WHERE email=$1`

	user := models.User{Email: email}

	err := u.client.QueryRow(ctx, q, email).Scan(&user.UUID, &user.Name, &user.Phone, &user.PasswordHash, &user.Crystals, &user.IsBlocked, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func NewUserCRUD(client postgres.Client) *UserCRUD {

	return &UserCRUD{
		client: client,
	}
}
