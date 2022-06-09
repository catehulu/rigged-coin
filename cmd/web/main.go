package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/catehulu/rigged-coin/internal/config"
	"github.com/catehulu/rigged-coin/internal/driver"
	"github.com/catehulu/rigged-coin/internal/handlers"
)

const portNumber string = ":8080"

var app config.AppConfig

func run() (*driver.DB, error) {

	log.Println("Connecting to datbase...")
	db, err := driver.ConnectMongoDB("mongodb://localhost:27017", "rigged")
	if err != nil {
		log.Fatal("Cannot connect to database")
		return nil, err
	}

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	return db, nil
}

// main is the main function
func main() {

	db, err := run()
	defer db.MongoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Staring application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
