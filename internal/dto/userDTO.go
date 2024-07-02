package dto

import "github.com/jackc/pgx/v5/pgtype"

type UserDTO struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Phone    string `json:"phone" binding:"required" example:"+78005553535"`
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password,omitempty" example:"UwU"`
}

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Phone    string `json:"phone" binding:"required" example:"+78005553535"`
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password,omitempty" example:"UwU"`
}

type UpdateUserDTO struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Phone string `json:"phone" binding:"required" example:"+78005553535"`
	Email string `json:"email" example:"john.doe@example.com"`
}

type ReadUserDTO struct {
	UUID      pgtype.UUID        `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name      string             `json:"name" example:"John Doe"`
	Phone     string             `json:"phone" example:"+78005553535"`
	Email     string             `json:"email" example:"john.doe@example.com"`
	Crystals  int                `json:"crystals" example:"12"`
	IsBlocked bool               `json:"is_blocked" example:"false"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at" example:"2020-01-01T00:00:00Z" swaggertype:"string"`
	CreatedAt pgtype.Timestamptz `json:"created_at" example:"2020-01-01T00:00:00Z" swaggertype:"string"`
}
