package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"FoundPlatform/model"
)

type DBStruct struct {
	Db *gorm.DB
}

var DB DBStruct

// ConnectDatabase connects to the database
func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgrespw dbname=postgres port=32768 sslmode=disable TimeZone=Asia/Singapore"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	log.Println("Running migrations...")
	// Migrate the schema
	err = db.AutoMigrate(&model.User{}, &model.Dataset{}, &model.Project{})
	if err != nil {
		panic(err)
	}

	log.Println("Migrations ran successfully")

	DB = DBStruct{
		Db: db,
	}
}
