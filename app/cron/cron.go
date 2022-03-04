// Package cron contains functions related to a timer or a ticker.
package cron

import (
	"github.com/frisk038/gocolormind/business/generate"
	"github.com/robfig/cron/v3"
)

func NewCron() *cron.Cron {
	c := cron.New()
	c.AddFunc("@midnight", generate.CreateFile)
	return c
}
