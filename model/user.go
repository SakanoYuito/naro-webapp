package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name	string	`json:"name" gorm:"not null;uniqueIndex"`
}

