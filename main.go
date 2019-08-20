package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dvonlehman/starter-api/app"
)

func main() {
	app := app.App{}
	app.Initialize()

	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Run(port))
}
