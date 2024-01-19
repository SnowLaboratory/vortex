package main

import (
	"snowlabs/vortex/controllers"
	"snowlabs/vortex/initializers"
	"snowlabs/vortex/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	r := gin.Default()
	r.HTMLRender = &TemplRender{}

	r.POST("/register", controllers.Register)
	r.GET("/register", controllers.RegisterPage)

	r.POST("/login", controllers.Login)
	r.GET("/login", controllers.LoginPage)

	r.GET("/dashboard", middleware.RequireAuth, controllers.Dashboard)

	r.Run()
}
