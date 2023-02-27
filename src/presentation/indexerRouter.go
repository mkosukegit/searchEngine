package presentation

import (
	"log"
	"net/http"
	"search/src/application/repository"
	"search/src/application/service/indexer"
)

type UserInfoRouter interface {
	Routing(w http.ResponseWriter, r *http.Request)
}

func Routing(w http.ResponseWriter, r *http.Request) {

	plainDocumentRepository := repository.NewPlainDocumentRepository()
	indexerService := indexer.NewUserInfoService(plainDocumentRepository)
	indexerController := NewIndexerController(indexerService)

	log.Print(r.Method)
	switch r.Method {
	case "GET":
		indexerController.createIndex(w, r)
	default:
		w.WriteHeader(405)
	}
}
