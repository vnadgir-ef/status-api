package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Get Status",
		Method:      "GET",
		Pattern:     "/status",
		HandlerFunc: StatusHandler,
	},
}
