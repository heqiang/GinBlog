package routers

import (
	"GinDemo/src/pkg/setting"
	"github.com/gin-gonic/gin"
)

func   InitRouter() *gin.Engine {
		r := gin.New()
		r.Use(gin.Recovery())
		r.Use(gin.Logger())
		gin.SetMode(setting.RunMode)
		r.GET("/Get", func(context *gin.Context) {
			context.JSON(200,gin.H{
				"Message":"test",
			})
		})
		return r
}