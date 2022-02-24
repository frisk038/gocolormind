package cron

import (
	"github.com/frisk038/gocolormind/business/generate"
	"github.com/robfig/cron/v3"
)

func StartNewCron() {
	c := cron.New()
	c.AddFunc("@midnight", generate.GenerateFile)
	c.Start()
}
