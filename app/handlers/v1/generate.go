package v1

import (
	"net/http"
	"time"

	"github.com/frisk038/gocolormind/business/generate"
	"github.com/gin-gonic/gin"
)

func Gen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	data, err := generate.ReadFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, combination{
		Date:        time.Now(),
		Combination: data,
	})
}