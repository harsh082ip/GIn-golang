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
	router.GET("/getquerystring", getQueryString)
	router.GET("/getparams/:name/:age", getParams)
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

	// reading body from Request
	body := c.Request.Body
	val, _ := io.ReadAll(body)

	// this is an object of response writer, gin.H is header
	c.JSON(200, gin.H{
		"data":      "Hello",
		"name":      "I am harsh",
		"body data": string(val),
	})
}

func getQueryString(c *gin.Context) {

	// reading query parameters
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"data": "This is getQueryString method",
		"name": name,
		"age":  age,
	})
}

func getParams(c *gin.Context) {

	name := c.Param("name")
	age := c.Param("age")

	// this is a object of response writer
	c.JSON(200, gin.H{
		"data": "This is getParams method",
		"name": name,
		"age":  age,
	})
}
