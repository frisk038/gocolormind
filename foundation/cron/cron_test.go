package cron

import (
	"os"
	"testing"

	"github.com/frisk038/gocolormind/business/generate"
)

func Test_task(t *testing.T) {
	t.Run("Succes", func(t *testing.T) {
		StartNewCron()
		_, err := os.Stat(generate.CombiFileName)
		if err != nil {
			t.Errorf("Combination File is not created ! (%s)", err)
		}
	})
}
