package repository

import (
	"FoundPlatform/database"
	"FoundPlatform/model"
)

// GetAllUsers gets all users from the database
func GetAllUsers() []model.User {
	var users []model.User
	db := database.DB.Db
	db.Find(&users)
	return users
}

// GetUserByID gets a user by id
func GetUserByID(id string) model.User {
	user := model.User{}
	db := database.DB.Db
	db.Where("ID = ?", id).First(&user)
	return user
}

// GetUserByEmail gets a user by email
func GetUserByEmail(email string) model.User {
	user := model.User{}
	db := database.DB.Db
	db.Where("Email = ?", email).First(&user)
	return user
}

// CreateUser creates a new user
func CreateUser(user model.User) model.User {
	db := database.DB.Db
	db.Create(&user)
	return user
}
