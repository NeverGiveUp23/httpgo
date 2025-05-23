package main

import (
	"flag"
	"fmt"
	"github.com/nevergiveup23/httpgo/internal/app"
	"github.com/nevergiveup23/httpgo/internal/routes"
	"net/http"
	"time"
)

func main() {

	var port int
	// using address of(&) operator to parse the value we are getting from the cammon into port integer variable port
	flag.IntVar(&port, "port", 8080, "go back and server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	} // panic == shutdown

	// defer at the end of execution
	defer app.DB.Close()

	// http.HandleFunc("/health", HealthCheck)
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are running our app on port %d!\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
