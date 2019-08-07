package tests

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../Models"
	"../Services"
)

const DB_URL = "mongodb://localhost:27017"

var id string

func setup() (srv Services.PeopleService, ctx context.Context, client *mongo.Client) {
	clientOptions := options.Client().ApplyURI(DB_URL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return
	}
	return Services.NewPeopleService(), context.Background(), client
}

func TestCreatePerson(t *testing.T) {
	srv, ctx, client := setup()
	person := Models.Person{Firstname: "Test person", Lastname: "Test lastName"}

	s, err := srv.CreatePerson(ctx, person, client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if s == "" {
		t.Errorf("Can't create a person in mongoddb")
	}
	id = s
}

func TestGetPerson(t *testing.T) {
	srv, ctx, client := setup()
	var person *Models.Person
	person, err := srv.GetPerson(ctx, "", client)
	if err == nil {
		t.Errorf("Error: %s", err)
	}
	person, err = srv.GetPerson(ctx, id, client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if person.Firstname == "" {
		t.Errorf("The database not return a person")
	}
}

func TestGetPeople(t *testing.T) {
	srv, ctx, client := setup()
	var people *[]Models.Person
	people, err := srv.GetPeople(ctx, client)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	flag := len(*people) != 0
	if !flag {
		t.Errorf("The database not return a list of people")
	}
}
