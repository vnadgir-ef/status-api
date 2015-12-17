package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	log.Printf("Starting server with these api endpoints...")
	for _, r := range routes {
		log.Printf("\t%s\t%s\n", r.Method, r.Pattern)
		router.Methods(r.Method).Path(r.Pattern).Name(r.Name).Handler(Logger(r.HandlerFunc, r.Name))
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./html/")))
	return router
}
