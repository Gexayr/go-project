package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);unique;index"`
	Password string `gorm:"type:varchar(255);not null"`
	Name     string `gorm:"type:varchar(255)"`
}
