package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// config struct to hold all the configuration settngs of the app
type config struct {
	port int
	env  string
}

// application struct to hold all the dependencies of the HTTP handlers,
// helpers, migrations.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Declare an instance of the config struct
	var cfg config

	// Read the value of the port and env command-line flags into the config struct.
	// We default to port number 4000 and the environment "development" if no flags are provided.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger which writes messagees to stdout stream,
	// prefixed with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Declare an instance of the application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare a HTTP server with some sensible timeout settings, which listents on
	// the port defined in the config struct and uses the servemux we created above as the handler.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server and log any errors that occur.
	logger.Printf("Starting %s server on %s\n", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
