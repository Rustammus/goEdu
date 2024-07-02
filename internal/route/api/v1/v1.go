package v1

import (
	"github.com/gin-gonic/gin"
	"goEdu/internal/service"
	"goEdu/pkg/auth"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(s *service.Services, t auth.TokenManager) *Handler {
	return &Handler{services: s, tokenManager: t}
}

func (h *Handler) Init(r *gin.RouterGroup) {
	h.initUsersRouter(r)
}
