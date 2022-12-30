package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CatFact struct {
	Fact   string `bson:"fact" json:"fact"`
	Length int    `bson:"length" json:"length"`
}

type Storer interface {
	GetAll() ([]*CatFact, error)
	Put(*CatFact) error
}

type MongoStore struct {
	client     *mongo.Client
	database   string
	collestion string
}

func NewMongoStore(uri string) (*MongoStore, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &MongoStore{
		client:     client,
		database:   "catfact",
		collestion: "facts",
	}, nil
}

func (store *MongoStore) Put(fact *CatFact) error {
	coll := store.client.Database(store.database).Collection(store.collestion)
	_, err := coll.InsertOne(context.TODO(), fact)
	return err
}

func (store *MongoStore) GetAll() ([]*CatFact, error) {
	coll := store.client.Database(store.database).Collection(store.collestion)

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	results := []*CatFact{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
