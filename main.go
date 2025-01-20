package main

import (
	"ratnesh2003/to-do-app/configs"
	"ratnesh2003/to-do-app/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnvVars()
	configs.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.RegisterTaskRoutes(r)

	r.Run()
}