package router

import (
	"net/http"

	"github.com/frodonLD/GoLangRESTAPIWithMux/handler"
	"github.com/gorilla/mux"
)

// Route is a custom struct that helps us to handle routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is the list of the routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Health",
		"GET",
		"/_health",
		handler.HealthCheckHandler,
	},
	Route{
		"GetAllNotifications",
		"GET",
		"/notifications",
		handler.GetAllNotificationsHandler,
	},
	Route{
		"GetNotification",
		"GET",
		"/notifications/{id}",
		handler.GetNotificationHandler,
	},
}

// NewRouter creates a new mux router with all routes of the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
