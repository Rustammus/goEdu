package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/dto"
)

type User struct {
	UUID   pgtype.UUID `json:"uuid"`
	Name   string      `json:"name"`
	Phone  string      `json:"phone"`
	Salary int         `json:"salary"`
}

func GetUserFromDTO(uuid pgtype.UUID, dto dto.UserDTO) *User {
	user := User{
		UUID:   uuid,
		Name:   dto.Name,
		Phone:  dto.Phone,
		Salary: dto.Salary,
	}
	return &user
}
