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
		"NmapScan",
		"POST",
		"/api/v1/nmapscan",
		NmapScan,
	},
	Route{
		"MasScan",
		"POST",
		"/api/v1/masscan",
		MasScan,
	},
}
