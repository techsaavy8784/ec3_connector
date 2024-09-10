package services

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
		"HealthCheck",
		"GET",
		"/api/rest/v1/healthcheck",
		HealthCheck,
	},
	/////////////////////////////////////////////////////////////////////////////////
	//                                FindResources                                //
	/////////////////////////////////////////////////////////////////////////////////

	Route{
		"FindResource",
		"GET",
		"/api/rest/v1/Find/{resource}",
		findResource,
	},
	Route{
		"GetResource",
		"GET",
		"/api/rest/v1/Get{resource}/{resId}",
		getResource,
	},
	Route{
		"SaveResource",
		"POST",
		"/{fhirSystem}/Patient",
		savePatient,
	},
	Route{
		"DebbieTest",
		"GET",
		"/{fhirSystem}/test",
		DebbieTest,
	},
	Route{
		"DebbieTest",
		"GET",
		"/{fhirSystem}/Patient/{id}",
		getPatient,
	},
}
