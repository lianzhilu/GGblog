package controller

import (
	errmsg "GGblog/internal/errormsg"
	"GGblog/internal/middleware/jwt"
	"GGblog/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var loginUser model.User
	ctx.ShouldBindJSON(&loginUser)

	var token string
	var code int
	var user model.User

	user, code = model.CheckLogin(loginUser.Username, loginUser.Password)

	if code == errmsg.SUCCESS {
		token, code = jwt.SetToken(user.Username)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"user":   user,
		"token":  token,
		"msg":    errmsg.GetErrorMessage(code),
	})
}
