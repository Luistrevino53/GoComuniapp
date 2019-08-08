package Endpoints

import (
	"context"
	"encoding/json"
	"net/http"

	"../Models"
)

type getPersonRequest struct {
	Id string `json:"_id"`
}

type getPersonResponse struct {
	Person Models.Person `json:"person"`
	Err    string        `json:"err,omitempty"`
}

type createPersonRequest struct {
	Person Models.Person `json:"person"`
}

type createPersonResponse struct {
	Id  string `json:"_id"`
	Err string `json:"err,omitempty"`
}

type getPeopleRequest struct{}

type getPeopleResponse struct {
	People []Models.Person `json:"people"`
	Err    string          `json:"err,omitempty"`
}

func DecodeCreatePersonRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createPersonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetPersonRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getPersonRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeGetPeopleRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getPeopleRequest
	return req, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
