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

func generate() []string {
	combi := []string{"black", "green", "brown", "yellow", "red", "purple"}
	return []string{combi[NewRandom.Intn(len(combi))], combi[NewRandom.Intn(len(combi))],
		combi[NewRandom.Intn(len(combi))], combi[NewRandom.Intn(len(combi))]}
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
