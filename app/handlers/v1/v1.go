// Package v1 regroup all handlers for v1 version.
package v1

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type Config struct {
	DB *sql.DB
}

func API(build string, log *log.Logger, cfg Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/combination", Gen)
	// router.POST("/cmbdb", cmbHand.Create)

	return router
}
