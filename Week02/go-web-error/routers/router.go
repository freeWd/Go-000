package routers

import (
	"go-gin-app/pkg/setting"
	v1 "go-gin-app/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/users", v1.GetUsers)
		apiv1.GET("/users/:id", v1.GetUserById)
	}
	return r
}
