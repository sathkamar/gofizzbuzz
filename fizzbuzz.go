package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to Go FizzBuzz",
			})
		})

		v1.GET("/getUserTODO", func(c *gin.Context) {

			userId := c.DefaultQuery("userId", "123123")

			c.JSON(http.StatusOK, gin.H{
				"message": "The TODO list will be sent for " + userId,
			})
		})

		v1.POST("/UpdateTODO", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "The Users TODO list will be sent",
			})
		})

		v1.PUT("/AddTODO", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "New TODO item added in users record",
			})
		})

		v1.DELETE("DeleteItem", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "The Users TODO Item deleted"})
		})
	}

	router.Run()
}
