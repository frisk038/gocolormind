// Package generate generates combination of the day for colormind.
package generate

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
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

const CombiFileName string = "/tmp/combination.txt"

var NewRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

func generate() []string {
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

func GenerateFile() {
	err := ioutil.WriteFile(CombiFileName,
		[]byte(strings.Join(generate(), ",")),
		0644)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFromFile() []string {
	data, err := ioutil.ReadFile(CombiFileName)
	switch err {
	case os.ErrNotExist:
		GenerateFile()
		data, err = ioutil.ReadFile(CombiFileName)
		if err != nil {
			fmt.Printf("Cant read combinantion file %s", err)
			return nil
		}
		return strings.Split(string(data), ",")
	case nil:
		return strings.Split(string(data), ",")
	default:
		fmt.Printf("Cant read combinantion file %s", err)
		return nil
	}
}
