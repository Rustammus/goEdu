package v1

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/crud"
	"goEdu/internal/dto"
	"goEdu/internal/models"
	"goEdu/internal/schemas"
)

func CreateUser(c *gin.Context) {
	var user dto.UserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		schemas.WriteResponseValidErr(c, err)
		return
	}

	userCRUD := crud.NewUserCRUD(crud.ConnPool)
	uuid, err := userCRUD.Create(context.TODO(), &user)
	if err != nil {
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "Data created correctly", uuid)
}

func ListAllUser(c *gin.Context) {
	userCRUD := crud.NewUserCRUD(crud.ConnPool)
	users, err := userCRUD.FindAll(context.TODO())
	if err != nil {
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "Data got correctly", users)
}

func FindUserByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	pgxUUID := pgtype.UUID{}
	err := pgxUUID.Scan(uuid)
	if err != nil {
		schemas.WriteResponseScanUUIDErr(c, err)
		return
	}

	userCRUD := crud.NewUserCRUD(crud.ConnPool)
	user, err := userCRUD.FindByUUID(context.TODO(), pgxUUID)
	if err != nil {
		if isEmptyRow := errors.Is(err, pgx.ErrNoRows); isEmptyRow {
			schemas.WriteResponseQueryEmptyRow(c, "User not find by UUID")
			return
		}
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "User got correctly", user)
}

func UpdateUserByUUID(c *gin.Context) {
	var userDTO dto.UserDTO
	uuid := c.Param("uuid")
	pgxUUID := pgtype.UUID{}

	if err := pgxUUID.Scan(uuid); err != nil {
		schemas.WriteResponseScanUUIDErr(c, err)
		return
	}

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		schemas.WriteResponseValidErr(c, err)
		return
	}
	user := models.GetUserFromDTO(pgxUUID, userDTO)
	userCRUD := crud.NewUserCRUD(crud.ConnPool)
	updatedUser, err := userCRUD.Update(context.TODO(), user)

	if err != nil {
		if isEmptyRow := errors.Is(err, pgx.ErrNoRows); isEmptyRow {
			schemas.WriteResponseQueryEmptyRow(c, "User not find by UUID")
			return
		}
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "User updated successfully", updatedUser)
}

func DeleteUserByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	pgxUUID := pgtype.UUID{}

	if err := pgxUUID.Scan(uuid); err != nil {
		schemas.WriteResponseScanUUIDErr(c, err)
		return
	}

	userCRUD := crud.NewUserCRUD(crud.ConnPool)
	err := userCRUD.Delete(context.TODO(), pgxUUID)
	if err != nil {
		if isEmptyRow := errors.Is(err, pgx.ErrNoRows); isEmptyRow {
			schemas.WriteResponseQueryEmptyRow(c, "User not find by UUID")
			return
		}
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "User deleted successfully", nil)
}
