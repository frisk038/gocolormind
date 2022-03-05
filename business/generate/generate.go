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

func Generate() string {
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

	return strings.Join(combi, ",")
}

func CreateFile() {
	data := Generate()
	f, err := os.Create(CombiFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wrt := bufio.NewWriter(f)
	_, err = wrt.WriteString(data)
	if err != nil {
		panic(err)
	}
	wrt.Flush()
}

func ReadFromFile() (string, error) {
	if _, err := os.Stat(CombiFileName); errors.Is(err, os.ErrNotExist) {
		CreateFile()
	}

	f, err := os.Open(CombiFileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	return scanner.Text(), nil
}
