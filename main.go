package main

import (
	"log"
	"os"
	"strconv"

	"github.com/dvonlehman/starter-api/app"
)

func main() {
	// var wait time.Duration
	// flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	// flag.Parse()

	app := app.App{}
	app.Initialize()

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// This function blocks unless an error was encountered
	err = app.Run(port)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server is running on port %d\n", port)

	// var srv *http.Server

	// // Run our server in a goroutine so that it doesn't block.
	// go func() {
	// 	srv, err = app.Run(port)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// c := make(chan os.Signal, 1)
	// // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	// signal.Notify(c, os.Interrupt)

	// // Block until we receive our signal.
	// <-c

	// // Create a deadline to wait for.
	// ctx, cancel := context.WithTimeout(context.Background(), wait)
	// defer cancel()
	// // Doesn't block if no connections, but will otherwise wait
	// // until the timeout deadline.
	// srv.Shutdown(ctx)
	// // Optionally, you could run srv.Shutdown in a goroutine and block on
	// // <-ctx.Done() if your application should wait for other services
	// // to finalize based on context cancellation.
	// log.Println("shutting down")
	// os.Exit(0)
}
