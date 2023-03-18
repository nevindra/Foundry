package repository

import (
	"FoundPlatform/database"
	"FoundPlatform/model"
)

func GetAllDatasets(id string) []model.Dataset {
	var datasets []model.Dataset
	db := database.DB.Db
	// Find all datasets for spesific user
	db.Where("UserID = ?", id).Find(&datasets)
	return datasets
}

func GetDatasetByID(id string) model.Dataset {
	dataset := model.Dataset{}
	db := database.DB.Db
	db.Where("ID = ?", id).First(&dataset)
	return dataset
}

func CreateDataset(dataset model.Dataset) model.Dataset {
	database.DB.Db.Create(&dataset)
	return dataset
}

func DeleteDataset(id string) {
	database.DB.Db.Where("ID = ?", id).Delete(&model.Dataset{})
}
