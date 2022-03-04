package main

import (
	"log"
	"os"

	v1 "github.com/frisk038/gocolormind/app/handlers/v1"
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/frisk038/gocolormind/app/cron"
)

func main() {
	build := "main"
	log := log.New(os.Stdout, "COLORMIND : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	port := os.Getenv("PORT")

	c := cron.NewCron()
	c.Start()

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	api := v1.API(build, log)

	api.Run(":" + port)
}
