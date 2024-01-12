package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	WEBPORT = ":8081"
)

func main() {

	router := gin.Default()

	router.GET("/getdata", func(c *gin.Context) {
		// this gin context contains every information about response
		// and request

		// this is an object of response writer, gin.H is header
		c.JSON(200, gin.H{
			"name":      "I am harsh",
			"Techstack": "golang",
		})
	})

	fmt.Printf("Starting server at port %v ...", WEBPORT)
	router.Run(WEBPORT)
}
