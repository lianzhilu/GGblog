package controller

import (
	errmsg "GGblog/internal/errormsg"
	"GGblog/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(ctx *gin.Context) {
	var article model.Article
	_ = ctx.ShouldBindJSON(&article)
	result := model.CreateArticle(&article)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  result,
		"article": article,
		"msg":     errmsg.GetErrorMessage(result),
	})
}

// 查询单个文章的详细信息
func GetArticleByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	article, code := model.GetArticleByID(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"article": article,
		"msg":     errmsg.GetErrorMessage(code),
	})
}

// 根据关键字搜索文章
func SearchArticle(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	keyword := ctx.Query("keyword")
	articles, code := model.GetArticlesByKeyWord(keyword, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status":   code,
		"articles": articles,
		"msg":      errmsg.GetErrorMessage(code),
	})
}

// 查询某个用户发布的所有文章
func SearchUserArticle(ctx *gin.Context) {
	userid, _ := strconv.Atoi(ctx.Param("userid"))
	articles, code := model.GetUserArticles(userid)

	ctx.JSON(http.StatusOK, gin.H{
		"status":   code,
		"articles": articles,
		"msg":      errmsg.GetErrorMessage(code),
	})
}

// 查询文章列表
func GetArticles(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	users, _ := model.GetArticles(pageSize, pageNum)
	code := errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"users":  users,
		"msg":    errmsg.GetErrorMessage(code),
	})
}

// 编辑文章
func EditArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var article model.Article
	ctx.ShouldBindJSON(&article)
	code := model.UpdateArticle(id, &article)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"article": article,
		"msg":     errmsg.GetErrorMessage(code),
	})
}

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	code := model.DeleteArticle(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrorMessage(code),
	})
}
