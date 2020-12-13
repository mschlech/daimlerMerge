package main

import (
	"net/http"

	"github.com/mschlech/daimlerMerge/cmd/handler"
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
		"healthCheck",
		"GET",
		"/healthcheck/",
		handler.GetHealthCheck,
	},
}
