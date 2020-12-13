package main

import (
	"net/http"

	"github.com/mschlech/daimler_merge/cmd/handler"
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
		"getMergeList",
		"get",
		"/mergedlist/",
		handler.GetMergeList,
	},
	Route{
		"healthCheck",
		"GET",
		"/healthcheck/",
		handler.GetHealthCheck,
	},
}
