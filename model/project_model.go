package model

// Project model is used to store the data of a project
// This model will have ID, Name and Description (optional)
// This model should have foreign key to the user who created it
type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	Datasets    []Dataset `gorm:"many2many:project_data;"`
}

// ProjectData is a table that stores which datasets are in which projects
// This table will have a foreign key to the project and a foreign key to the dataset
type ProjectData struct {
	ID        string  `json:"id"`
	ProjectID Project `json:"project_id" gorm:"foreignKey:ProjectID;references:ID"`
	DatasetID Dataset `json:"dataset_id" gorm:"foreignKey:DatasetID;references:ID"`
}
