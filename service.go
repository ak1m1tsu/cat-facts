package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// CatFactWorker is an interface that collects facts about cats from
// https://catfact.ninja
type CatFactWorker interface {
	Start() error
}

// catFactWorker implements the CatFactWorker interface
type catFactWorker struct {
	endpoint string
	store    Storer
}

// NewCatFactWorker makes new instance of CatFactWorker interface with
// MongoDB Client and Cat Facts API endpoint
func NewCatFactWorker(store Storer, endpoint string) CatFactWorker {
	return &catFactWorker{
		store:    store,
		endpoint: endpoint,
	}
}

// Start every 2 seconds makes request to API endpoint and
// collects response data to MongoDB
func (cfw *catFactWorker) Start() error {
	ticker := time.NewTicker(2 * time.Second)
	for {
		resp, err := http.Get(cfw.endpoint)
		if err != nil {
			return err
		}

		var catFact CatFact
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}

		if err := cfw.store.Put(&catFact); err != nil {
			return err
		}

		<-ticker.C
	}
}
