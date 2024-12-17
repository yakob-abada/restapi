package main

import (
	"github.com/yakob-abada/restapi/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Apply fixtures
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Match{}, &model.Profile{})

	err = db.Create(getUsers()).Error
	if err != nil {
		panic(err)
	}

	err = db.Create(getProfiles()).Error
	if err != nil {
		panic(err)
	}

	err = db.Create(getExplore()).Error
	if err != nil {
		panic(err)
	}
}

func getUsers() []*model.Match {
	return []*model.Match{
		{RecipientUserId: 1, ActorUserId: 3, Status: 2, CreatedAt: time.Now()},
		{RecipientUserId: 1, ActorUserId: 4, Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: 1, ActorUserId: 5, Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: 1, ActorUserId: 6, Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: 1, ActorUserId: 7, Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: 1, ActorUserId: 8, Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: 1, ActorUserId: 9, Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: 1, ActorUserId: 10, Status: 0, CreatedAt: time.Now()},
		{RecipientUserId: 1, ActorUserId: 11, Status: 0, CreatedAt: time.Now()},
	}
}

func getProfiles() []*model.Profile {
	return []*model.Profile{
		{ID: 1, Name: "John Doe", Gender: "Male", Lat: "0", Long: "0", DietType: "vegan", Age: 25},
		{ID: 2, Name: "Sara Doe", Gender: "Female", Lat: "0", Long: "1", DietType: "vegan", Age: 26},
		{ID: 3, Name: "Jakob Doe", Gender: "Male", Lat: "0", Long: "2", DietType: "vegan", Age: 29},
		{ID: 4, Name: "Steve Doe", Gender: "Male", Lat: "0", Long: "3", DietType: "vegan", Age: 20},
		{ID: 5, Name: "Tom Doe", Gender: "Male", Lat: "0", Long: "4", DietType: "vegetarian", Age: 22},
		{ID: 6, Name: "Ola Doe", Gender: "Female", Lat: "0", Long: "5", DietType: "vegan", Age: 35},
		{ID: 7, Name: "Ali Doe", Gender: "Male", Lat: "0", Long: "6", DietType: "vegan", Age: 23},
		{ID: 8, Name: "Koko Doe", Gender: "Male", Lat: "0", Long: "7", DietType: "vegan", Age: 26},
		{ID: 9, Name: "Tota Doe", Gender: "Female", Lat: "0", Long: "8", DietType: "vegetarian", Age: 25},
		{ID: 10, Name: "Hanks Doe", Gender: "Male", Lat: "0", Long: "5", DietType: "vegan", Age: 27},
		{ID: 11, Name: "Suzi Doe", Gender: "Female", Lat: "0", Long: "2", DietType: "vegan", Age: 40},
	}
}

func getExplore() []*model.Explore {
	return []*model.Explore{
		{ID: 1, DietType: "vegan", AgeFrom: 20, AgeTo: 30, DistanceRadius: 111195, Gender: "female"},
	}
}
