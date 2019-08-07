package Endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"../Models"
)

type getPersonRequest struct {
	id string `json:"_id"`
}

type getPersonResponse struct {
	person *Models.Person `json:"person"`
	Err    string         `json:"err,omitempty"`
}

type createPersonRequest struct {
	person Models.Person `json:"person"`
}

type createPersonResponse struct {
	id  string `json:"_id"`
	Err string `json:"err,omitempty"`
}

type getPeopleRequest struct{}

type getPeopleResponse struct {
	people []Models.Person `json:"people"`
	Err    string          `json:"err,omitempty"`
}

func decodeCreatePersonRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createPersonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetPersonRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getPersonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetPeopleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getPeopleRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
