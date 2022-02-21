// Package v1 regroup all handlers for v1 version.
package v1

import (
	"log"
	"net/http"

	"github.com/frisk038/gocolormind/business/generate"
	"github.com/gin-gonic/gin"
)

func API(build string, log *log.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, generate.Generate())
	})

	return router
}
