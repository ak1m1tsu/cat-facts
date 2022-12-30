package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	store Storer
}

// NewServe makes new instance of Server
func NewServer(s Storer) *Server {
	return &Server{
		store: s,
	}
}

// handleGetAllFacts returns json response that contains all facts in MongoDB Client
func (server *Server) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	catFacts, err := server.store.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Context-Type", "application/json")
	json.NewEncoder(w).Encode(catFacts)
}
