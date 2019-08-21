package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dvonlehman/starter-api/app"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	app := app.App{}
	app.Initialize()

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run(port))
}
