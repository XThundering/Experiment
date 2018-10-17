package main

import (
	"github.com/XThundering/Experiment/plugin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": plugin.Print(),
		})
	})
	r.Run()
}
