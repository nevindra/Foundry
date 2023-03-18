package main

import (
	"FoundPlatform/database"
	"FoundPlatform/routes"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	r := gin.Default()
	GetRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
	r.Use(gin.Logger())

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	return r
}

func GetRoutes(r *gin.Engine) {
	routes.AddUserRoutes(r)
	routes.AddDatasetRoutes(r)
}

func main() {
	database.ConnectDatabase()
	SetupServer()
}
