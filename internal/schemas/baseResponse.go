package schemas

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData any

type BaseResponse struct {
	Massage string       `json:"massage"`
	Data    ResponseData `json:"data,omitempty"`
	Error   string       `json:"error,omitempty"`
}

func WriteResponseValidErr(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, BaseResponse{
		Massage: "Validation error",
		Error:   err.Error(),
	})
}

func WriteResponseQueryErr(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, BaseResponse{
		Massage: "Creating user error",
		Error:   err.Error(),
	})
}

func WriteResponseQueryEmptyRow(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, BaseResponse{
		Massage: msg,
	})
}

func WriteResponseScanUUIDErr(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, BaseResponse{
		Massage: "Scan UUID error",
		Error:   err.Error(),
	})
}

func WriteResponseOk(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, BaseResponse{
		Massage: msg,
		Data:    data,
	})
}
