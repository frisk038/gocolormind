// Package v1 regroup all handlers for v1 version.
package v1

import (
	"database/sql"
	"log"

	"github.com/frisk038/gocolormind/business/data/combination"
	"github.com/gin-gonic/gin"
)

func API(build string, log *log.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())

	cmbHand := Handlers{
		Combi: combination.NewCore(&sql.DB{}),
	}

	router.GET("/combination", Gen)
	router.POST("/newCombinationDB", cmbHand.Create)

	return router
}
