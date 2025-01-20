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
		"task": task,
	})
}

func TaskComplete(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	configs.DB.First(&task, id)

	configs.DB.Model(&task).Updates(models.Task{
		IsCompleted: true,
	})

	c.JSON(200, gin.H{
		"task": task,
	})
}

func TaskUpdate(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	configs.DB.First(&task, id)

	var body struct {
		Title       string
		Description string
		Deadline	time.Time
	}

	c.Bind(&body)

	configs.DB.Model(&task).Updates(models.Task{
		Title: body.Title,
		Description: body.Description,
		Deadline: body.Deadline,
	})

	c.JSON(200, gin.H{
		"task": task,
	})
}

func TaskAll(c *gin.Context) {
	var tasks []models.Task

	configs.DB.Find(&tasks)

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func TaskDelete(c *gin.Context) {
	id := c.Param("id")
	configs.DB.Delete(&models.Task{}, id)

	c.Status(200)
}