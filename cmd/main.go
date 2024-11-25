package main

import (
	"Cloud-Native-Go-KVStorage/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.HelloWorld)

	r.HandleFunc("/v1/{key}", handler.KeyValuePutHandler).Methods(http.MethodPut)
	r.HandleFunc("/v1/{key}", handler.KeyValueGetHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
