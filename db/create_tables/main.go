package main

import (
	"github.com/yakob-abada/restapi/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Apply migration
	err = db.AutoMigrate(&model.Match{}, &model.Profile{}, &model.Explore{})
	if err != nil {
		panic(err)
	}
}
