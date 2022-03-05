// Package combination provides an example of a core business API. Right now
// these calls are just wrapping the data/store layer. But at some point you
// will want auditing or something that isn't specific to the data/store layer.
package combination

import (
	"database/sql"
	"time"

	"github.com/frisk038/gocolormind/business/data/db"
	"github.com/frisk038/gocolormind/business/generate"
	"github.com/gin-gonic/gin"
)

type DailyCombination struct {
	Combi     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Core struct {
	store db.Store
}

// NewCore constructs a core for product api access.
func NewCore(str *sql.DB) Core {
	return Core{
		store: db.NewStore(str),
	}
}

func (cr Core) Create(c *gin.Context, now time.Time) (DailyCombination, error) {
	dbCmb := db.DailyCombination{
		Combi:     generate.Generate(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if _, err := cr.store.AddCombination(dbCmb); err != nil {
		return DailyCombination{}, err
	}

	return DailyCombination(dbCmb), nil
}

// func toCmb(dbCmb db.DailyCombination) DailyCombination {
// 	cmb := (*DailyCombination)(unsafe.Pointer(&dbCmb))
// 	return *cmb
// }
