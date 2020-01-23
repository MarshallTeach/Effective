package routers

import (
	_ "Effective/docs"
	"Effective/pkg/setting"
	"Effective/routers/api/auth"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())


	gin.SetMode(setting.ServerSetting.RunMode)


	apiauth := r.Group("/api/auth")
	{
		apiauth.POST("/register", auth.Register)
		apiauth.POST("/login", auth.Login)
	}
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
