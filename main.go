package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"./Endpoints"
	"./Services"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := Services.NewPeopleService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := Endpoints.PeopleEndpoints{
		GetPersonEndpoint:    Endpoints.MakeGetPerson(srv),
		GetPeopleEndpoint:    Endpoints.MakeGetPeople(srv),
		CreatePersonEndpoint: Endpoints.MakeCreatePerson(srv),
	}

	go func() {
		log.Println("the service is listening on port: ", *httpAddr)
		handler := NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()
	log.Fatalln(<-errChan)
}
