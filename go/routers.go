/*
 * Swagger Petstore
 *
 * This is a sample Petstore server.  You can find  out more about Swagger at  [http://swagger.io](http://swagger.io) or on  [irc.freenode.net, #swagger](http://swagger.io/irc/).
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	// LINE API
	Route{
		"Index",
		"POST",
		"/message/receive",
		MessageReceive,
	},
	{
		"",
		"POST",
		"/room/add",
		RoomAdd,
	},
	{
		"",
		"GET",
		"/rooms",
		ListRoom,
	},
	{
		"",
		"GET",
		"/rooms/type",
		ListAllRoomTypes,
	},
	{
		"",
		"POST",
		"/rooms/search",
		SearchRoom,
	},

	// Schedule
	{
		"",
		"POST",
		"/schedule/add",
		AddSchedule,
	},
	{
		"",
		"GET",
		"/schedules",
		ListSchedule,
	},
	{
		"",
		"POST",
		"/schedules/filter",
		FilterSchedule,
	},

	Route{
		"Index",
		"GET",
		"/kp1ay/test/firebase/",
		Hello,
	},

	Route{
		"Index",
		"GET",
		"/kp1ay/test/1.0.0/",
		Index,
	},
}
