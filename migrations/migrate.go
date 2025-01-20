package main

import (
	"ratnesh2003/to-do-app/configs"
	"ratnesh2003/to-do-app/models"
)

func init() {
	configs.LoadEnvVars()
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Task{})
}
