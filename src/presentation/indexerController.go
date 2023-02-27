package presentation

import (
	"context"
	"encoding/json"
	"net/http"
	"search/src/application/service/indexer"
)

type IndexerController interface {
	createIndex(w http.ResponseWriter, r *http.Request)
}

type indexerController struct {
	indexerService indexer.IndexerService
}

func NewIndexerController(indexerService indexer.IndexerService) IndexerController {
	return &indexerController{indexerService}
}

func (indexerController *indexerController) createIndex(w http.ResponseWriter, r *http.Request) {

	// var userReq domain.User
	// json.NewDecoder(r.Body).Decode(&userReq)

	err := indexerController.indexerService.CreateIndex(context.Background())
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("users")
}
