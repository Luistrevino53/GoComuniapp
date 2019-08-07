package Services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"

	"../Models"
)

const (
	DB_NAME         = "testdb"
	COLLECTION_NAME = "people"
)

var client *mongo.Client

type PeopleService interface {
	CreatePerson(ctx context.Context, person Models.Person, client *mongo.Client) (string, error)
	GetPerson(ctx context.Context, id string, client *mongo.Client) (*Models.Person, error)
	GetPeople(ctx context.Context, client *mongo.Client) (*[]Models.Person, error)
}

type dataPeopleService struct{}

func NewPeopleService() PeopleService {
	return dataPeopleService{}
}

func (dataPeopleService) CreatePerson(ctx context.Context, person Models.Person, client *mongo.Client) (string, error) {
	collection := client.Database(DB_NAME).Collection(COLLECTION_NAME)
	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		return "error", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (dataPeopleService) GetPerson(ctx context.Context, id string, client *mongo.Client) (*Models.Person, error) {
	var person Models.Person
	if id == "" {
		return &person, errors.New("id parameter is nil or blank")
	}
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &person, err
	}

	collection := client.Database(DB_NAME).Collection(COLLECTION_NAME)
	err = collection.FindOne(ctx, Models.Person{ID: hexId}).Decode(&person)
	if err != nil {
		return &person, err
	}
	return &person, nil
}

func (dataPeopleService) GetPeople(ctx context.Context, client *mongo.Client) (*[]Models.Person, error) {
	var people []Models.Person
	collection := client.Database(DB_NAME).Collection(COLLECTION_NAME)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return &people, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Models.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		return &people, err
	}
	return &people, nil
}
