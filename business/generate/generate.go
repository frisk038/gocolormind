// Package generate generates combination of the day for colormind.
package generate

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const CombiFileName string = "/tmp/combination.txt"

var NewRandom = rand.New(rand.NewSource(time.Now().UnixNano()))

// Why loop if we can use a map?
func colorInCombi(color string, combi []string) bool {
	for _, b := range combi {
		if color == b {
			return true
		}
	}
	return false
}

func generate() []string {
	combiLst := []string{"black", "green", "brown", "yellow", "red", "purple"}
	combi := make([]string, 4)
	for i := range combi {
		color := combiLst[NewRandom.Intn(len(combiLst))]
		if colorInCombi(color, combi) {
			color = combiLst[NewRandom.Intn(len(combiLst))]
		}
		combi[i] = color
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
	switch {
	case errors.Is(err, os.ErrNotExist):
		GenerateFile()
		data, err = ioutil.ReadFile(CombiFileName)
		if err != nil {
			return []string{fmt.Sprintf("Cant read generated combination file %s", err)}
		}
		return strings.Split(string(data), ",")
	case err == nil:
		return strings.Split(string(data), ",")
	default:
		return []string{fmt.Sprintf("Cant read combination file %s", err)}
	}
}
