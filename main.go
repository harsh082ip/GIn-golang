package main

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

const (
	WEBPORT = ":8081"
)

func main() {

	router := gin.Default()

	router.GET("/getdata", getData)
	router.POST("/postdata", postData)

	fmt.Printf("Starting server at port %v ...", WEBPORT)
	router.Run(WEBPORT)
}

func getData(c *gin.Context) {
	// this gin context contains every information about response
	// and request

	// this is an object of response writer, gin.H is header
	c.JSON(200, gin.H{
		"name":      "I am harsh",
		"Techstack": "golang",
	})
}

func postData(c *gin.Context) {
	// this gin context contains every information about response
	// and request

	body := c.Request.Body
	val, _ := io.ReadAll(body)

	// this is an object of response writer, gin.H is header
	c.JSON(200, gin.H{
		"data":      "Hello",
		"name":      "I am harsh",
		"body data": string(val),
	})
}
