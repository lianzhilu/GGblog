package model

import (
	errmsg "GGblog/internal/errormsg"
	"fmt"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string `gorm:"type: varchar(100);not null" json:"title"`
	Description string `gorm:"type: varchar(200)" json:"description"`
	Content     string `gorm:"type: longtext;not null" json:"content"`
	UserID      uint   `gorm:"not null" json:"userid"`
	User        User   `gorm:"foreignkey:UserID"`
}

// 添加文章
func CreateArticle(article *Article) int {
	if len(article.Content) == 0 {
		return errmsg.ERROR_ART_IS_EMPTY
	}

	err := db.Create(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单个文章的详细信息
func GetArticleByID(id int) (Article, int) {
	var article Article
	err := db.Preload("User").Where("id = ?", id).First(&article).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return article, errmsg.ERROR
	}
	if err == gorm.ErrRecordNotFound {
		return article, errmsg.ERROR_ART_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// 根据关键字搜索文章
func GetArticlesByKeyWord(keyword string, pageSize int, pageNum int) ([]Article, int) {
	if len(keyword) == 0 {
		return nil, errmsg.ERROR_KW_IS_EMPTY
	}
	keyword = "%" + keyword + "%"

	var articles []Article
	if pageNum != -1 {
		pageNum = (pageNum - 1) * pageSize
	}

	err := db.Preload("User").Where("title LIKE ?", keyword).Or("description LIKE ?", keyword).Or("content LIKE ?", keyword).Limit(pageSize).Offset(pageNum).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	if err == gorm.ErrRecordNotFound || len(articles) == 0 {
		return nil, errmsg.ERROR_RESULT_NOT_FOUND
	}

	return articles, errmsg.SUCCESS
}

// 查询某个用户发布的所有文章
func GetUserArticles(userid int) ([]Article, int) {
	var articles []Article
	err := db.Where("user_id = ?", userid).Find(&articles).Error
	fmt.Println(err)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	if err == gorm.ErrRecordNotFound || len(articles) == 0 {
		return nil, errmsg.ERROR_RESULT_NOT_FOUND
	}
	return articles, errmsg.SUCCESS
}

// 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, error) {
	var articles []Article
	if pageNum != -1 {
		pageNum = (pageNum - 1) * pageSize
	}

	err := db.Preload("User").Limit(pageSize).Offset(pageNum).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// 编辑文章
func UpdateArticle(id int, article *Article) int {
	updateMap := make(map[string]interface{})
	updateMap["title"] = article.Title
	updateMap["description"] = article.Description
	updateMap["content"] = article.Content
	err := db.Model(&Article{}).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
