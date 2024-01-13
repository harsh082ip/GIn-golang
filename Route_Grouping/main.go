package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	WEBPORT = ":8081"
)

func main() {

	router := gin.Default()

	router.GET("/getquerystring", getQueryString)
	router.GET("/getparams/:name/:age", getParams)

	auth := gin.BasicAuth(gin.Accounts{
		"admin": "admin123",
	})
	admin := router.Group("/admin", auth)
	{
		admin.GET("/getdata", getData)
	}

	client := router.Group("/client")
	{
		client.POST("/postdata", postData)
	}

	// default http configuration
	server := &http.Server{
		Addr:         WEBPORT,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
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
