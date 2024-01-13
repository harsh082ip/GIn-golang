package main

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/harsh082ip/Gin-golang/gin_middleware/Middleware"
)

func main() {

	router := gin.Default()
	// router.Use(middleware.Authenticate)  apply to all routes

	admin := router.Group("/admin", middleware.Authenticate) // apply to a group
	{

		admin.GET("/getData3", getData3)
	}

	router.GET("/getData1", middleware.Authenticate, getData1)                       // apply to a specific route
	router.GET("/getData2", middleware.Authenticate, middleware.AddHeader, getData2) // chaning middlewares

	router.Run(":8082")
}

func getData1(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "Hi I am getData1 Method",
	})
}

func getData2(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "Hi I am getData2 Method",
	})
}

func getData3(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "Hi I am getData3 Method",
	})
}
