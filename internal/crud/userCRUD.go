package crud

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/dto"
	"goEdu/internal/models"
	"goEdu/internal/repos"
	"goEdu/pkg/client/postgres"
)

type UserCRUD struct {
	client postgres.Client
}

func (u *UserCRUD) Create(ctx context.Context, userDTO *dto.UserDTO) (pgtype.UUID, error) {

	q := `INSERT INTO public.users (name, phone, salary) VALUES ($1, $2, $3) RETURNING uuid`
	uuid := pgtype.UUID{}
	err := u.client.QueryRow(ctx, q, userDTO.Name, userDTO.Phone, userDTO.Salary).Scan(&uuid)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return uuid, nil
}

func (u *UserCRUD) FindAll(ctx context.Context) ([]models.User, error) {

	q := `SELECT uuid, name, phone, salary FROM public.users`
	rows, err := u.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0)

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.UUID, &user.Name, &user.Phone, &user.Salary)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

func (u *UserCRUD) FindByUUID(ctx context.Context, uuid pgtype.UUID) (*models.User, error) {

	q := `SELECT uuid, name, phone, salary FROM public.users WHERE uuid=$1`

	user := models.User{}

	err := u.client.QueryRow(ctx, q, uuid).Scan(&user.UUID, &user.Name, &user.Phone, &user.Salary)
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (u *UserCRUD) Update(ctx context.Context, user *models.User) (*models.User, error) {
	//TODO implement me
	q := `
		UPDATE public.users
		SET (name, phone, salary) = ($2, $3, $4)
		WHERE uuid = $1 RETURNING uuid, name, phone, salary`
	err := u.client.QueryRow(context.TODO(), q, user.UUID, user.Name, user.Phone, user.Salary).
		Scan(&user.UUID, &user.Name, &user.Phone, &user.Salary)
	if err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (u *UserCRUD) Delete(ctx context.Context, uuid pgtype.UUID) error {
	//TODO implement me
	q := `DELETE FROM public.users WHERE uuid = $1 RETURNING uuid`
	err := u.client.QueryRow(context.TODO(), q, uuid).Scan(&uuid)
	if err != nil {
		return err
	}
	return nil
}

func NewUserCRUD(client postgres.Client) repos.UserRepository {

	return &UserCRUD{
		client: client,
	}
}
