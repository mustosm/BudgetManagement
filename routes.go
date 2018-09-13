package main

import (
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

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

var routes = Routes{
	Route{
		"GetTransactions",
		"GET",
		"/transactions",
		GetTransactions,
	},
	Route{
		"HeadTransactions",
		"HEAD",
		"/transactions",
		GetTransactions,
	},
	Route{
		"GetTransactionById",
		"GET",
		"/transactions/{transactionId}",
		GetTransactionByID,
	},
	Route{
		"HeadTransactionsById",
		"HEAD",
		"/transactions/{transactionId}",
		GetTransactions,
	},
	Route{
		"PostTransaction",
		"POST",
		"/transactions",
		PostTransaction,
	},
	Route{
		"PutTransaction",
		"PUT",
		"/transactions/{transactionId}",
		PutTransaction,
	},
	Route{
		"GetHealthCheck",
		"GET",
		"/health",
		GetHealthCheck,
	},
}
