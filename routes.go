package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Ping",
		"GET",
		"/api/v1/ping",
		Ping,
	},
	Route{
		"DocumentIndex",
		"GET",
		"/api/v1/documents",
		DocumentIndex,
	},
	Route{
		"DocumentCreate",
		"POST",
		"/api/v1/documents",
		DocumentCreate,
	},
	Route{
		"AuthKeyCreate",
		"POST",
		"/authkeys",
		AuthKeyCreate,
	},
	Route{
		"AuthKeyDelete",
		"DELETE",
		"/authkeys/{authKeysId}",
		AuthKeyCreate,
	},
}
