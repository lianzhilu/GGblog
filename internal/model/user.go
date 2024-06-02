package model

import (
	errmsg "GGblog/internal/errormsg"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(100); not null" json:"username"`
	Password string `gorm:"type: varchar(100); not null" json:"password"`
	Gender   int    `gorm:"type: int" json:"gender"`
}

// 用bcrypt对密码进行加密
func EncryptPw(password string) string {
	hashPw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashPw)
}

// 查询用户是否存在
func CheckUser(username string) int {
	fmt.Println(username)
	var user User
	// result := db.Where("username = ?", username).First(&user)
	// result := db.First(&user, "username = ?", username)
	result := db.Where("username = ?", username).First(&user)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 通过钩子函数在创建用户前先对密码加密
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Password = EncryptPw(u.Password)
	return nil
}

// 添加用户
func CreateUser(user *User) int {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, error) {
	var users []User
	if pageNum != -1 {
		pageNum = (pageNum - 1) * pageSize
	}

	err := db.Limit(pageSize).Offset(pageNum).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 通过关键字查询用户
func GetUsersByKeyWord(keyword string, pageSize int, pageNum int) ([]User, int) {
	if len(keyword) == 0 {
		return nil, errmsg.ERROR_KW_IS_EMPTY
	}
	keyword = "%" + keyword + "%"

	var users []User
	if pageNum != -1 {
		pageNum = (pageNum - 1) * pageSize
	}

	err := db.Where("username LIKE ?", keyword).Limit(pageSize).Offset(pageNum).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	if err == gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR_RESULT_NOT_FOUND
	}

	return users, errmsg.SUCCESS
}

// 编辑用户信息
func UpdateUser(id int, user *User) int {
	updateMap := make(map[string]interface{})
	updateMap["username"] = user.Username
	updateMap["gender"] = user.Gender
	err := db.Model(&User{}).Where("id = ?", id).Updates(updateMap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var emptyUser User
	err := db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return emptyUser, errmsg.ERROR_USER_NOT_EXIST
	}

	pwErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if pwErr != nil {
		return emptyUser, errmsg.ERROR_PASSWORD_WRONG
	}

	return user, errmsg.SUCCESS
}

// 将所有密码加密
func EncryptAllPw() {
	var users []User
	db.Find(&users)
	for i := 0; i < len(users); i++ {
		id := users[i].ID
		password := users[i].Password
		hashPw := EncryptPw(password)

		updateMap := make(map[string]interface{})
		updateMap["password"] = hashPw

		db.Model(&User{}).Where("id = ?", id).Updates(updateMap)
	}
}

// 修改密码
// func ChangePassword(id int, oldPW string, newPW string) {
// 	var user User
// 	err := db.Where("id = ?", id).First(&user)
// 	err
// }
