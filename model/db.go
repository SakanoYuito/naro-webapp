package model

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "user:password@tcp(db)/app?charset=utf8mb4&parseTime=True&loc=Local"
	time.Sleep(3 * time.Second)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}
}