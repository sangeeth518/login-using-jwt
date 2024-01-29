package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sangeeth/jwt-go/controllers"
	"github.com/sangeeth/jwt-go/initializers"
	"github.com/sangeeth/jwt-go/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/signout", controllers.AdminSignout, middleware.Userauth)
	r.Run()
}
