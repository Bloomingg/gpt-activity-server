package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname string `gorm:"not null"`
	Work     string
	Phone    string
	Activity uint `gorm:"index;not null"`
}
