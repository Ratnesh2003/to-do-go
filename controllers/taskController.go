package controllers

import (
	"ratnesh2003/to-do-app/configs"
	"ratnesh2003/to-do-app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		Deadline	time.Time
	}
	c.Bind(&body)

	task := models.Task{Title: body.Title, Description: body.Description, Deadline: body.Deadline}

	result := configs.DB.Create(&task)

	if result.Error != nil {
		c.Status(500)
	}

	c.JSON(201, gin.H{
		"message": task,
	})
}
