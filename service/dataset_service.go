package service

import (
	"FoundPlatform/model"
	"FoundPlatform/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

// GetAllDatasets returns all datasets for a user
func GetAllDatasets(c *gin.Context) {
	id := c.Param("user_id")
	datasets := repository.GetAllDatasets(id)
	// check if datasets is found
	if len(datasets) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Dataset not found"})
		return
	}
	c.JSON(http.StatusOK, datasets)
}

// GetDatasetByID returns a dataset by id
func GetDatasetByID(c *gin.Context) {
	id := c.Param("id")
	dataset := repository.GetDatasetByID(id)

	// check if dataset is found
	if dataset.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Dataset not found"})
		return
	}

	c.JSON(http.StatusOK, dataset)
}

// CreateDataset uploads a new file dataset
func CreateDataset(c *gin.Context) {
	dataset := model.Dataset{}

	// Get value from form
	dataset.Name = c.PostForm("name")
	dataset.Description = c.PostForm("description")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)

	// check if file is csv
	if filepath.Ext(filename) != ".csv" {
		c.String(http.StatusBadRequest, "file is not csv")
		return
	}

	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	// Save the dataset to the database
	dataset.Url = filename
	repository.CreateDataset(dataset)
	c.JSON(http.StatusOK, dataset)
}

// DeleteDataset deletes a dataset by id
// This function will also delete the file from the server
func DeleteDataset(c *gin.Context) {
	id := c.Param("id")

	// check if dataset exists
	dataset := repository.GetDatasetByID(id)
	if dataset.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Dataset not found"})
		return
	}

	// delete the file from the server
	path := dataset.Url
	err := os.RemoveAll(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	// delete the dataset from the database
	repository.DeleteDataset(id)

	c.JSON(http.StatusOK, gin.H{"message": "Dataset deleted"})
}
