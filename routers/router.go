package routers

import (
	"GGblog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppConf.Mode)
	r := gin.Default()

	router := r.Group("/api/v1")
	{
		router.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}

	r.Run(utils.AppConf.Port)
}
