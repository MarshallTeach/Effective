package routers

import (
	_ "Effective/docs"
	"Effective/routers/api/v1/basic"
	"Effective/routers/api/v1/content"

	/*"Effective/middleware/jwt"*/
	"Effective/pkg/setting"
	"Effective/routers/api/auth"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)

	apiv1 := r.Group("/api/v1")
	/*apiv1.Use(jwt.JWT())*/
	{
		apiv1.GET("/basicInfo", basic.GetBasicInfo)
		apiv1.POST("/basicInfo", basic.EditBasicInfo)
		apiv1.GET("/column", content.GetColumn)
	}

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
