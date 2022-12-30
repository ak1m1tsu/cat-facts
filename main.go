package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mongoPort := flag.String("mp", "27017", "The mongodb port.")
	servicePort := flag.String("p", "3000", "The microservice port.")
	flag.Parse()

	storage, err := NewMongoStore(fmt.Sprintf("mongodb://localhost:%s", *mongoPort))
	if err != nil {
		log.Fatal(err)
	}

	worker := NewCatFactWorker(storage, "https://catfact.ninja/fact")
	go worker.Start()

	server := NewServer(storage)
	http.HandleFunc("/facts", server.handleGetAllFacts)
	http.ListenAndServe(fmt.Sprintf(":%s", *servicePort), nil)
}
