package v1

import (
	"net/http"
	"time"

	"github.com/frisk038/gocolormind/business/generate"
	"github.com/gin-gonic/gin"
)

type CombinPld struct {
	Date        time.Time `json:"date"`
	Combination []string  `json:"combi"`
}

// Handlers manages the set of product endpoints.
// type Handlers struct {
// 	Combi
// }

// func (h Handlers) Create(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	newCombi, err := h.Combi.Create(c, time.Now())
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	c.IndentedJSON(http.StatusCreated, CombinPld{
// 		Date: newCombi.CreatedAt,
// 		// Combination: newCombi.Combi,
// 	})
// }

func Gen(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	data, err := generate.ReadFromFile()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, CombinPld{
		Date:        time.Now(),
		Combination: data,
	})
}
