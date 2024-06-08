package controller

import (
	errmsg "GGblog/internal/errormsg"
	"GGblog/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"url":    url,
		"msg":    errmsg.GetErrorMessage(code),
	})
}
