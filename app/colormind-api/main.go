package main

import (
	"database/sql"
	"log"
	"os"

	v1 "github.com/frisk038/gocolormind/app/handlers/v1"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"

	"github.com/frisk038/gocolormind/app/cron"
)

func main() {
	build := "main"
	log := log.New(os.Stdout, "COLORMIND : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	port := os.Getenv("PORT")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	c := cron.NewCron()
	c.Start()

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	cfg := v1.Config{
		DB: db,
	}

	api := v1.API(build, log, cfg)

	api.Run(":" + port)
}
