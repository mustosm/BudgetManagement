package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// GetTransationByID methode provide the fonctionnality to Get a transaction by Id
func GetTransationByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	amount := Amount{Value: 9.5, Currency: "EUR"}
	transaction := Transaction{TransactionID: "0001", Description: "First transaction", Amount: amount, TransactionType: "transaction", PaymentMeans: "Cash", CategoryID: "0001", AccountID: "0001", Month: "Mars", Year: "2018", CreationDate: time.Now(), UpdateDate: time.Now()}

	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		panic(err)
	}

}

// GetTransations methode provide the fonctionnality to  Get a list of transactions
func GetTransations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	transactions := Transactions{
		Transaction{TransactionID: "1234", Description: "First transaction"},
		Transaction{TransactionID: "1235", Description: "Second transaction"},
		Transaction{TransactionID: "1236", Description: "Second transaction"},
	}

	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		panic(err)
	}

}
