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

	// Apply drop table.
	if db.Migrator().HasTable(&model.Match{}) {
		err = db.Migrator().DropTable(&model.Match{})
		if err != nil {
			panic(err)
		}
	}

	if db.Migrator().HasTable(&model.Profile{}) {
		err = db.Migrator().DropTable(&model.Profile{})
		if err != nil {
			panic(err)
		}
	}

	if db.Migrator().HasTable(&model.Explore{}) {
		err = db.Migrator().DropTable(&model.Explore{})
		if err != nil {
			panic(err)
		}
	}
}
