package main

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"./Endpoints"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints.PeopleEndpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/people").Handler(httptransport.NewServer(
		endpoints.GetPeopleEndpoint,
		Endpoints.DecodeGetPeopleRequest,
		Endpoints.EncodeResponse))

	r.Methods("GET").Path("/people/{id}").Handler(httptransport.NewServer(
		endpoints.GetPersonEndpoint,
		Endpoints.DecodeGetPersonRequest,
		Endpoints.EncodeResponse,
	))

	r.Methods("POST").Path("/people").Handler(httptransport.NewServer(
		endpoints.CreatePersonEndpoint,
		Endpoints.DecodeCreatePersonRequest,
		Endpoints.EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
