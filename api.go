package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	client *mongo.Client
}

// NewServe makes new instance of Server
func NewServer(client *mongo.Client) *Server {
	return &Server{
		client: client,
	}
}

// handleGetAllFacts returns json response that contains all facts in MongoDB Client
func (server *Server) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	coll := server.client.Database("catfact").Collection("facts")

	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	results := []bson.M{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Context-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
