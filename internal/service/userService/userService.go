package userService

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/dto"
	"goEdu/internal/repos"
	"goEdu/internal/schemas/requestSchemas"
	"goEdu/pkg/auth"
	"goEdu/pkg/hash"
	"time"
)

type UserService struct {
	Repo         repos.UserRepository
	TokenManager auth.TokenManager
	Hasher       hash.PasswordHasher
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *UserService) Create(dto *dto.CreateUserDTO) (pgtype.UUID, error) {
	pswHash, err := u.Hasher.Hash(dto.Password)
	if err != nil {
		return pgtype.UUID{}, err
	}
	dto.Password = pswHash
	uuid, err := u.Repo.Create(context.TODO(), dto)
	return uuid, err
}

func (u *UserService) ListAll() ([]dto.ReadUserDTO, error) {
	users, err := u.Repo.FindAll(context.TODO())
	return users, err
}

func (u *UserService) FindByUUID(uuid pgtype.UUID) (*dto.ReadUserDTO, error) {
	user, err := u.Repo.FindByUUID(context.TODO(), uuid)
	return user, err
}

func (u *UserService) UpdateByUUID(uuid pgtype.UUID, userDTO *dto.UpdateUserDTO) (*dto.ReadUserDTO, error) {
	user, err := u.Repo.Update(context.TODO(), userDTO, uuid)
	return user, err
}

func (u *UserService) DeleteByUUID(uuid pgtype.UUID) error {
	err := u.Repo.Delete(context.TODO(), uuid)
	return err
}

func (u *UserService) LogIn(userSignIn *requestSchemas.InputUserSignIn) (Tokens, error) {
	pswHash, err := u.Hasher.Hash(userSignIn.Password)
	if err != nil {
		return Tokens{}, err
	}

	user, err := u.Repo.FindByEmailWithHash(context.TODO(), userSignIn.Email)
	if err != nil {
		return Tokens{}, err
	}
	if pswHash == user.PasswordHash {
		return u.createTokens(user.UUID)
	}
	return Tokens{}, errors.New("Wrong password")
}

func (u *UserService) createTokens(uuid pgtype.UUID) (Tokens, error) {
	t := Tokens{}
	uuidString := fmt.Sprintf("%x", uuid.Bytes)
	token, err := u.TokenManager.NewJWT(uuidString, time.Minute*20)
	if err != nil {
		return t, err
	}

	refToken, err := u.TokenManager.NewJWT(uuidString, time.Hour*12)
	if err != nil {
		return t, err
	}

	t.AccessToken = token
	t.RefreshToken = refToken

	return t, nil
}
