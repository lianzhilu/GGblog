package controller

import (
	errmsg "GGblog/internal/errormsg"
	"GGblog/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserIsExist(ctx *gin.Context) {

}

// 添加用户
func AddUser(ctx *gin.Context) {
	var user model.User
	_ = ctx.ShouldBindJSON(&user)
	fmt.Println(user.Username)

	code := model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
		// fmt.Println("not in database")
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"user":   user,
		"msg":    errmsg.GetErrorMessage(code),
	})
}

// 查询用户列表
func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	users, _ := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"users":  users,
		"msg":    errmsg.GetErrorMessage(code),
	})
}

// 通过关键字查询用户
func SearchUser(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	keyword := ctx.Query("keyword")
	users, code := model.GetUsersByKeyWord(keyword, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"users":  users,
		"msg":    errmsg.GetErrorMessage(code),
	})
}

// 编辑用户
func EditUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var user model.User
	ctx.ShouldBindJSON(&user)
	code := model.CheckUser(user.Username)

	if code == errmsg.SUCCESS {
		code = model.UpdateUser(id, &user)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		ctx.Abort()
	}

	// code = model.UpdateUser(id, &user)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"user":   user,
		"msg":    errmsg.GetErrorMessage(code),
	})
}

// 删除用户
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrorMessage(code),
	})
}

// 将原有的数据的明文储存的代码加密
func UpdateAllPw(ctx *gin.Context) {
	model.EncryptAllPw()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
