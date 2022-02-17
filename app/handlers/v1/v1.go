// Package v1 regroup all handlers for v1 version.
package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/frisk038/gocolormind/business/generate"
)

func API(build string, log *log.Logger) *httptreemux.ContextMux {
	tm := httptreemux.NewContextMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		gen := struct {
			Combi []string
		}{
			Combi: generate.Generate(),
		}
		json.NewEncoder(w).Encode(gen)
	}

	tm.Handle(http.MethodGet, "/combination", h)

	return tm
}
