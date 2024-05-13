package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"type: varchar(100);not null" json:"title"`
	Content string `gorm:"type: longtext;not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"userid"`
	User    User   `gorm:"foreignkey:UserID"`
}
