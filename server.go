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
		Endpoints.decodeGetPeopleRequest,
		Endpoints.encodeResponse))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
