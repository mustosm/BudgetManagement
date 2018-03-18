package main

import (
	"encoding/json"
	"net/http"
)

func GetTransationById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	transaction := Transaction{transactionId: "1234", description: "First transaction"}
	
	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		panic(err)
	}
}

func GetTransations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	transactions := Transactions{
		Transaction{transactionId: "1234", description: "First transaction"},
		Transaction{transactionId: "1235", description: "Second transaction"},
	}

	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		panic(err)
	}
}
