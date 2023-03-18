package routes

import (
	"FoundPlatform/service"
	"github.com/gin-gonic/gin"
)

// AddDatasetRoutes adds dataset routes to the router
func AddDatasetRoutes(r *gin.Engine) {
	dataset := r.Group("/dataset")
	{
		dataset.GET("/user/:user_id", service.GetAllDatasets)
		dataset.GET("/:id", service.GetDatasetByID)
		dataset.POST("/create", service.CreateDataset)
		dataset.DELETE("/:id", service.DeleteDataset)
	}
}
