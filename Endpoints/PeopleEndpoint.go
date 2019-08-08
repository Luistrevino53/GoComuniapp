package Endpoints

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../Models"
	"../Services"
)

const DB_URL = "mongodb://localhost:27017"

type PeopleEndpoints struct {
	GetPersonEndpoint    endpoint.Endpoint
	GetPeopleEndpoint    endpoint.Endpoint
	CreatePersonEndpoint endpoint.Endpoint
}

func MakeGetPerson(srv Services.PeopleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		client := setup(ctx)
		req := request.(getPersonRequest)
		s, err := srv.GetPerson(ctx, req.id, client)
		if err != nil {
			return getPersonResponse{s, err.Error()}, nil
		}
		return getPersonResponse{s, ""}, nil
	}
}

func MakeCreatePerson(srv Services.PeopleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		client := setup(ctx)
		req := request.(createPersonRequest)
		p, err := srv.CreatePerson(ctx, req.person, client)
		if err != nil {
			return createPersonResponse{p, err.Error()}, nil
		}
		return createPersonResponse{p, ""}, nil
	}
}

func MakeGetPeople(srv Services.PeopleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		client := setup(ctx)
		_ = request.(getPeopleRequest)
		list, err := srv.GetPeople(ctx, client)
		if err != nil {
			return getPeopleResponse{list, err.Error()}, nil
		}
		return getPeopleResponse{list, ""}, nil
	}
}

func (e PeopleEndpoints) GetPerson(ctx context.Context) (*Models.Person, error) {
	req := getPersonRequest{}
	res, err := e.GetPersonEndpoint(ctx, req)
	if err != nil {
		return new(Models.Person), err
	}
	getResp := res.(getPersonResponse)
	if getResp.Err != "" {
		return new(Models.Person), errors.New(getResp.Err)
	}
	return getResp.person, nil

}

func (e PeopleEndpoints) GetPeople(ctx context.Context) ([]Models.Person, error) {
	req := getPeopleRequest{}
	res, err := e.GetPeopleEndpoint(ctx, req)
	if err != nil {
		return *new([]Models.Person), err
	}
	peopleResp := res.(getPeopleResponse)
	if peopleResp.Err != "" {
		return *new([]Models.Person), errors.New(peopleResp.Err)
	}
	return peopleResp.people, nil
}

func setup(ctx context.Context) *mongo.Client {
	clientOptions := options.Client().ApplyURI(DB_URL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil
	}
	return client
}
