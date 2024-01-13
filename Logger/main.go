package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harsh082ip/Gin-golang/Logger/logger_middleware"
	"github.com/mattn/go-colorable"
)

func main() {

	router := gin.Default()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint formatted information in %v %v %v %v \n", absolutePath, httpMethod, handlerName, nuHandlers)
	}

	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	f, _ := os.Create("ginLogging.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // writing logs to this file and with stdout

	router.Use(gin.LoggerWithFormatter(logger_middleware.FormatLogsJson))

	router.GET("/getdata", getData)
	router.Run(":8082")
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi I am GetData Method",
	})
}
