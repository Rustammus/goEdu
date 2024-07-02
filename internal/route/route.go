package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "goEdu/internal/route/api/v1"
	"goEdu/internal/service"
	"goEdu/pkg/auth"
)

type Handler struct {
	Services     *service.Services
	TokenManager auth.TokenManager
}

func NewHandler(s *service.Services, t auth.TokenManager) *Handler {
	return &Handler{Services: s, TokenManager: t}
}

func (h *Handler) Init() *gin.Engine {
	r := gin.Default()
	handlerV1 := v1.NewHandler(h.Services, h.TokenManager)
	apiV1 := r.Group("/api/v1")
	handlerV1.Init(apiV1)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
