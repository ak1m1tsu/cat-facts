package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CatFactWorker is an interface that collects facts about cats from
// https://catfact.ninja
type CatFactWorker interface {
	Start() error
}

// catFactWorker implements the CatFactWorker interface
type catFactWorker struct {
	endpoint string
	client   *mongo.Client
}

// NewCatFactWorker makes new instance of CatFactWorker interface with
// MongoDB Client and Cat Facts API endpoint
func NewCatFactWorker(client *mongo.Client, endpoint string) CatFactWorker {
	return &catFactWorker{
		client:   client,
		endpoint: endpoint,
	}
}

// Start every 2 seconds makes request to API endpoint and
// collects response data to MongoDB
func (cfw *catFactWorker) Start() error {
	coll := cfw.client.Database("catfact").Collection("facts")
	ticker := time.NewTicker(2 * time.Second)

	for {
		resp, err := http.Get(cfw.endpoint)
		if err != nil {
			return err
		}
		var catFact bson.M
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}
		_, err = coll.InsertOne(context.TODO(), catFact)
		if err != nil {
			return err
		}
		<-ticker.C
	}
}
