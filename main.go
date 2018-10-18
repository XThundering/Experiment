package main

import (
	"github.com/XThundering/Experiment/plugin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"return": plugin.Print(),
		})
	})
	r.Run()
}
