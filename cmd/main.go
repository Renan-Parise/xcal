package main

import (
	"github.com/gin-gonic/gin"
	"github.com/renan-parise/gofreela/internal/repositories"
	"github.com/renan-parise/gofreela/router"
)

func main() {
	db := repositories.Connect()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "alive",
		})
	})

	router.JobsRouter(r, db)

	r.Run("127.0.0.1:8000")
}
