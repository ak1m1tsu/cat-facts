package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoPort := flag.String("mp", "27017", "The mongodb port.")
	servicePort := flag.String("p", "3000", "The microservice port.")
	flag.Parse()
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%s", *mongoPort)),
	)
	if err != nil {
		panic(err)
	}

	worker := NewCatFactWorker(client, "https://catfact.ninja/fact")
	go worker.Start()

	server := NewServer(client)
	http.HandleFunc("/facts", server.handleGetAllFacts)
	http.ListenAndServe(fmt.Sprintf(":%s", *servicePort), nil)
}
