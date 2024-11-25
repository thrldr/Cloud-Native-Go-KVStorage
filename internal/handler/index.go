package handler

import (
	"Cloud-Native-Go-KVStorage/internal/storage"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	if _, err := w.Write([]byte("hello world\n")); err != nil {
		log.Fatalf("error handling request: %v", err)
	}
}

// @Route /v1/key/{key}
func KeyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := storage.Put(key, string(value)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

// @Route /v1/key/{key}
func KeyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := storage.Put(key, string(value)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
