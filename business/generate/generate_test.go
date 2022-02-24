// Package generate generates combination of the day for colormind.
package generate

import (
	"os"
	"testing"
)

func isCombiFileExists(t *testing.T) {
	_, err := os.Stat(CombiFileName)
	if err != nil {
		t.Errorf("Combination File is not created ! (%s)", err)
	}
}

func TestGenerateFile(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		GenerateFile()
		isCombiFileExists(t)
	})
	t.Run("already_exists", func(t *testing.T) {
		GenerateFile()
		isCombiFileExists(t)
		GenerateFile()
		isCombiFileExists(t)
	})
}

func TestReadFromFile(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		GenerateFile()
		isCombiFileExists(t)
		arr := ReadFromFile()
		switch {
		case len(arr) != 4:
			t.Error("arr len in file is not 4 ")
		case arr == nil:
			t.Error("arr is nil")
		default:
			return
		}
	})
	t.Run("file_deleted", func(t *testing.T) {
		os.Remove(CombiFileName)
		arr := ReadFromFile()
		if len(arr) != 4 {
			t.Error("arr len in file is not 4 ")
		}
	})
}
