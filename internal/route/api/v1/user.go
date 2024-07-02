package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/dto"
	"goEdu/internal/schemas"
	"goEdu/internal/schemas/requestSchemas"
)

func (h *Handler) initUsersRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("", h.userCreate)
		user.GET("/list", h.userListAll)
		user.GET("/:uuid", h.userFindByUUID)
		user.POST("login", h.userLogIn)
		authenticated := user.Group("", h.userIdentity)
		{
			authenticated.PUT("/:uuid", h.userUpdateByUUID)
			authenticated.DELETE("/:uuid", h.userDeleteByUUID)
		}

	}
}

// CreateUser godoc
// @Summary      Create user
// @Description  Create users
// @Tags         User API
// @Accept       json
// @Produce      json
// @Param CreateUserDTO body dto.CreateUserDTO true "User base"
// @Success 200 {object} schemas.BaseResponse
// @Failure      500  {object}	schemas.BaseResponse
// @Router       /user [post]
func (h *Handler) userCreate(c *gin.Context) {
	var user dto.CreateUserDTO

	if err := c.ShouldBindJSON(&user); err != nil {
		schemas.WriteResponseValidErr(c, err)
		return
	}

	uuid, err := h.services.User.Create(&user)
	if err != nil {
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "Data created correctly", uuid)
}

// ListAllUser godoc
// @Summary      List users
// @Description  List all users
// @Tags         User API
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.ReadUserDTO
// @Failure      500  {object}	schemas.BaseResponse
// @Router       /user/list [get]
func (h *Handler) userListAll(c *gin.Context) {
	users, err := h.services.User.ListAll()
	if err != nil {
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "Data got correctly", users)
}

// FindUserByUUID godoc
// @Summary      Find User by uuid
// @Description  Find User by uuid
// @Tags         User API
// @Accept       json
// @Produce      json
// @Param uuid path string true "User UUID" format(uuid)
// @Success 200 {object} dto.ReadUserDTO
// @Failure      500  {object}	schemas.BaseResponse
// @Router       /user/{uuid} [get]
func (h *Handler) userFindByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	pgxUUID := pgtype.UUID{}
	err := pgxUUID.Scan(uuid)
	if err != nil {
		schemas.WriteResponseScanUUIDErr(c, err)
		return
	}

	user, err := h.services.User.FindByUUID(pgxUUID)
	if err != nil {
		if isEmptyRow := errors.Is(err, pgx.ErrNoRows); isEmptyRow {
			schemas.WriteResponseQueryEmptyRow(c, "User not found by UUID")
			return
		}
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "User got correctly", user)
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Update users
// @Tags         User API
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param CreateUserDTO body dto.UpdateUserDTO true "User base"
// @Param uuid path string false "User UUID" format(uuid)
// @Success 200 {object} dto.ReadUserDTO
// @Failure      500  {object}	schemas.BaseResponse
// @Router       /user/{uuid} [put]
func (h *Handler) userUpdateByUUID(c *gin.Context) {
	var userDTO dto.UpdateUserDTO
	uuid, ok := c.Get(userCtx)
	if !ok {
		schemas.WriteResponseScanUUIDErr(c, errors.New("uuid in user context is empty"))
		return
	}
	pgxUUID := pgtype.UUID{}
	if err := pgxUUID.Scan(uuid); err != nil {
		schemas.WriteResponseScanUUIDErr(c, err)
		return
	}

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		schemas.WriteResponseValidErr(c, err)
		return
	}

	updatedUser, err := h.services.User.UpdateByUUID(pgxUUID, &userDTO)
	if err != nil {
		if isEmptyRow := errors.Is(err, pgx.ErrNoRows); isEmptyRow {
			schemas.WriteResponseQueryEmptyRow(c, "User not found by UUID")
			return
		}
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "User updated successfully", updatedUser)
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Delete users
// @Tags         User API
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param uuid path string false "User UUID" format(uuid)
// @Success 200 {object} schemas.BaseResponse
// @Failure      500  {object}	schemas.BaseResponse
// @Failure      400  {object}	schemas.BaseResponse
// @Router       /user/{uuid} [delete]
func (h *Handler) userDeleteByUUID(c *gin.Context) {

	uuid, ok := c.Get(userCtx)
	if !ok {
		schemas.WriteResponseScanUUIDErr(c, errors.New("uuid in user context is empty"))
		return
	}

	pgxUUID := pgtype.UUID{}
	if err := pgxUUID.Scan(uuid); err != nil {
		schemas.WriteResponseScanUUIDErr(c, err)
		return
	}

	err := h.services.User.DeleteByUUID(pgxUUID)
	if err != nil {
		if isEmptyRow := errors.Is(err, pgx.ErrNoRows); isEmptyRow {
			schemas.WriteResponseQueryEmptyRow(c, "User not found by UUID")
			return
		}
		schemas.WriteResponseQueryErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "User deleted successfully", nil)
}

// SignInUser godoc
// @Summary      LogIn user
// @Description  LogIn users
// @Tags         UserAuth API
// @Accept       json
// @Produce      json
// @Param InputUserSignIn body requestSchemas.InputUserSignIn true "User email and password"
// @Success 200 {object} schemas.BaseResponse
// @Failure      500  {object}	schemas.BaseResponse
// @Router       /user/login [post]
func (h *Handler) userLogIn(c *gin.Context) {

	userSignIn := requestSchemas.InputUserSignIn{}

	if err := c.ShouldBindJSON(&userSignIn); err != nil {
		schemas.WriteResponseValidErr(c, err)
		return
	}

	tokens, err := h.services.User.LogIn(&userSignIn)
	if err != nil {
		schemas.WriteResponseValidErr(c, err)
		return
	}

	schemas.WriteResponseOk(c, "Authenticated", tokens)
}
