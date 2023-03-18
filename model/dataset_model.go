package model

import "gorm.io/gorm"

// Dataset model is used to store the data of a dataset
// This model will have ID, Name, Description, and Url
// This model should have foreign key to the user who created it
type Dataset struct {
	gorm.Model
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	UserID      string    `json:"user_id"`
	Projects    []Project `gorm:"many2many:project_data;"`
}
