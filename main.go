package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yakob-abada/restapi/bootstrap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()

	router := gin.Default()
	router.GET("/profile_explore", bootstrap.NewProfileHandler(db).Explore)
	router.GET("/match", bootstrap.NewMatchHandler(db).WeMatched)
	router.GET("/like", bootstrap.NewLikeHandler(db).WhoLikedMe)

	err := router.Run("localhost:8080")
	if err != nil {
		panic("Could not start server: " + err.Error())
	}
}

func initDB() *gorm.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open db error: " + err.Error())
	}

	return db
}
