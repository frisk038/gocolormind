// Package generate generates combination of the day for colormind.
package generate

import (
	"math/rand"
	"time"
)

const (
	black = iota + 1
	green
	brown
	yellow
	red
	purple
)

var NewRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

func Generate() []string {
	coulours := map[int]string{black: "black", green: "green", brown: "brown",
		yellow: "yellow", red: "red", purple: "purple"}
	combi := make([]string, 4)

	for i := range combi {
		random := NewRandom.Intn(5) + 1
		coulour := coulours[random]
		combi[i] = coulour
	}

	return combi
}
