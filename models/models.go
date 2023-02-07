package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
}

type Todo struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(100);null"`
	Completed   bool
	UserID      uint `gorm:"not null"`
}
