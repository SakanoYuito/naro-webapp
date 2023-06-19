package model

import (
	// "time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name	string	`json:"name" gorm:"not null;default:null"`
}

