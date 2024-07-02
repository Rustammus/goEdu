package route

import (
	"github.com/gin-gonic/gin"
	"goEdu/internal/config"
	"goEdu/internal/route/api/v1"
)

func SetupRouter(conf config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/hello", v1.Greetings)
		userAPI := apiv1.Group("/user")
		{
			userAPI.POST("", v1.CreateUser)
			userAPI.GET("/list", v1.ListAllUser)
			userAPI.GET("/:uuid", v1.FindUserByUUID)
			userAPI.PUT("/:uuid", v1.UpdateUserByUUID)
			userAPI.DELETE("/:uuid", v1.DeleteUserByUUID)
		}
	}
	return r
}
