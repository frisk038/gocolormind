package v1

import (
	"net/http"
	"time"

	"github.com/frisk038/gocolormind/business/data/combination"
	"github.com/frisk038/gocolormind/business/generate"
	"github.com/gin-gonic/gin"
)

type CombinPld struct {
	Date        time.Time `json:"date"`
	Combination string    `json:"combi"`
}

// Handlers manages the set of product endpoints.
type Handlers struct {
	Combi combination.Core
}

func (h Handlers) Create(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	now := time.Now()
	newCombi, err := h.Combi.Create(c, now)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, CombinPld{
		Date:        now,
		Combination: newCombi.Combi,
	})
}

func Gen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	data, err := generate.ReadFromFile()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, CombinPld{
		Date:        time.Now(),
		Combination: data,
	})
}
