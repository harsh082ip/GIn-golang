package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	fmt.Println(c.Request.Header)
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Token not present",
		})
	}
	c.Next()
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key", "value")
	c.Next()
}
