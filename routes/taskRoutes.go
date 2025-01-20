package routes

import (
	"ratnesh2003/to-do-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	router.POST("/task/", controllers.TaskCreate)
}