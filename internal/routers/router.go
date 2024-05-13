package routers

import (
	"GGblog/internal/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(setting.AppConf.Mode)
	r := gin.Default()

	router := r.Group("/api/v1")
	{
		router.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}

	r.Run(setting.AppConf.Port)
}
