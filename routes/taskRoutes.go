package routes

import (
	"ratnesh2003/to-do-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(router *gin.Engine) {
	router.POST("/task/", controllers.TaskCreate)
	router.PUT("/task/complete/:id", controllers.TaskComplete)
	router.PUT("/task/:id", controllers.TaskUpdate)
	router.GET("/task/all", controllers.TaskAll)
	router.DELETE("/task/:id", controllers.TaskDelete)
}