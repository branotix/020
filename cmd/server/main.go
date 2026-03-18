package main

import (
	"github.com/branotix/p2p/config"
	"github.com/branotix/p2p/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.ConnectDatabase()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
