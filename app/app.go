package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.Path("/api/time").HandlerFunc(handleTime)
	router.Path("/api/time").Queries("tz", "{tz}").HandlerFunc(handleTime)

	log.Fatal(http.ListenAndServe("localhost:9090", router))
}
