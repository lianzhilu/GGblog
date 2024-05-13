package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(100); not null" json:"username"`
	Password string `gorm:"type: varchar(100); not null" json:"password"`
	Gender   int    `gorm:"type: int" json:"gender"`
}
