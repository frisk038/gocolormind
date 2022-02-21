// Package v1 regroup all handlers for v1 version.
package v1

import (
	"log"
	"net/http"
	"time"

	"github.com/frisk038/gocolormind/business/generate"
	"github.com/gin-gonic/gin"
)

type combination struct {
	Date        time.Time `json:"date"`
	Combination []string  `json:"combi"`
}

func API(build string, log *log.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	gin.SetMode(gin.ReleaseMode)

	router.GET("/combination", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, combination{
			Date:        time.Now(),
			Combination: generate.Generate(),
		})
	})

	return router
}
