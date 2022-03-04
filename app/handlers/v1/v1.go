// Package v1 regroup all handlers for v1 version.
package v1

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type combination struct {
	Date        time.Time `json:"date"`
	Combination []string  `json:"combi"`
}

func API(build string, log *log.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/combination", Gen)

	return router
}
