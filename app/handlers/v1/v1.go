// Package v1 regroup all handlers for v1 version.
package v1

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func API(build string, log *log.Logger) *gin.Engine {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/combination", func(c *gin.Context) {
		c.String(http.StatusOK, "generate.Generate()", nil)
	})

	router.Run(":" + port)

	return router
}
