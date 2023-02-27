package web

import (
	"net/http"
	"search/src/presentation"
)

type Middleware interface {
	Connect()
}

func Connect() {
	// http.HandleFunc("/api/user/register/", controller.Routing)
	http.HandleFunc("/api/inverted_index/", presentation.Routing)
}
