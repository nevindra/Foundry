package routes

import (
	"FoundPlatform/service"
	"github.com/gin-gonic/gin"
)

// AddUserRoutes adds user routes to the router
func AddUserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/", service.GetALlUsers)
		user.GET("/:id", service.GetUserByID)
		user.POST("/create", service.CreateUser)
	}
}
