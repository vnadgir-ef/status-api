package main

import (
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

func Logger(inner http.Handler, name string) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, inner)
}
