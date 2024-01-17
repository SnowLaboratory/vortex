package main

import (
	"snowlabs/vortex/controllers"
	"snowlabs/vortex/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.Run()
}
