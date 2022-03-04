// Package generate generates combination of the day for colormind.
package generate

import (
	"bufio"
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

const CombiFileName string = "/tmp/combination.txt"

func colorInCombi(color string, combi []string) bool {
	for _, b := range combi {
		if color == b {
			return true
		}
	}
	return false
}

func generate() []string {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	combiLst := []string{"black", "green", "brown", "yellow", "red", "purple"}
	combi := make([]string, 4)
	for i := range combi {
		color := combiLst[random.Intn(len(combiLst))]
		if colorInCombi(color, combi) {
			color = combiLst[random.Intn(len(combiLst))]
		}
		combi[i] = color
	}
	return combi
}

func CreateFile() {
	data := []byte(strings.Join(generate(), ","))
	err := os.WriteFile(CombiFileName, data, 0644)
	if err != nil {
		panic(err)
	}
}

func ReadFromFile() ([]string, error) {
	if _, err := os.Stat(CombiFileName); errors.Is(err, os.ErrNotExist) {
		CreateFile()
	}

	f, err := os.Open(CombiFileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	return strings.Split(scanner.Text(), ","), nil
}
