package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// GetTransactionByID methode provide the fonctionnality to Get a transaction by Id
func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	var err error
	var transaction Transaction
	params := mux.Vars(r)
	fields := r.URL.Query().Get("fields")
	if fields != "" {
		transaction, err = dao.FindSelectById(params["transactionId"], strings.Split(fields, ","))
	} else {
		transaction, err = dao.FindById(params["transactionId"])
	}
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respond(w, http.StatusOK, transaction)
}

// GetTransactions methode provide the fonctionnality to Get a list of transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	var err error
	var transactions Transactions
	fields := r.URL.Query().Get("fields")
	if fields != "" {
		transactions, err = dao.FindSelect(strings.Split(fields, ","))
	} else {
		transactions, err = dao.FindAll()
	}
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respond(w, http.StatusOK, transactions)
}

// PostTransaction methode provide the fonctionnality to save a transaction
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	tid := bson.NewObjectId()
	transaction.TransactionID = &tid
	if transaction.TransactionDate.IsZero() {
		tdate := time.Now().UTC()
		transaction.TransactionDate = &tdate
	}
	cdate := time.Now().UTC()
	udate := time.Now().UTC()
	transaction.CreationDate = &cdate
	transaction.UpdateDate = &udate
	transaction.Month = transaction.UpdateDate.Month().String()
	transaction.Year = strconv.Itoa(transaction.UpdateDate.Year())

	if err := dao.Insert(transaction); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond(w, http.StatusCreated, nil)
}

// PutTransaction methode provide the fonctionnality to update a transaction
func PutTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	udate := time.Now().UTC()
	transaction.UpdateDate = &udate

	if err := dao.Update(transaction, params["transactionId"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respond(w, http.StatusCreated, nil)
}

// Health check methode provide the service status
func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := dao.HealthCheck()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var health HealthCheck
	health.Message = "I'm alive"
	respond(w, http.StatusOK, health)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	jsonErr := JsonErr{Code: code, Message: msg}
	respond(w, code, jsonErr)
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error when marshaling payload")
			panic(err)
		}
	}
}
