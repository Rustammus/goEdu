package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	UUID         pgtype.UUID        `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name         string             `json:"name" example:"John"`
	Phone        string             `json:"phone" example:"+78005553535"`
	Email        string             `json:"email" example:"john.doe@example.com"`
	PasswordHash string             `json:"-" example:"UwU"`
	Crystals     int                `json:"crystals" example:"12"`
	Tasks        []Task             `json:"tasks,omitempty" example:"[tasks]"`
	IsBlocked    bool               `json:"is_blocked" example:"false"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at" example:"2020-01-01T00:00:00Z"`
	CreatedAt    pgtype.Timestamptz `json:"created_at" example:"2020-01-01T00:00:00Z"`
}
