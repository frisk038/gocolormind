// Package v1 regroup all handlers for v1 version.
package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

func API(build string, log *log.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/combination", Gen)
	router.POST("/newCombinationDB")

	return router
}
