package main

import (
	"time"
)

type Transaction struct {
	transactionId   string    `json:"transactionId,omitempty"`
	description     string    `json:"description,omitempty"`
	amount          Amount    `json:"amount,omitempty"`
	transactionType string    `json:"transactionType,omitempty"`
	paymentMeans    string    `json:"paymentMeans,omitempty"`
	categoryId      string    `json:"categoryId,omitempty"`
	accountId       string    `json:"accountId,omitempty"`
	month           string    `json:"month,omitempty"`
	year            string    `json:"year,omitempty"`
	creationDate    time.Time `json:"creationDate,omitempty"`
	updateDate      time.Time `json:"updateDate,omitempty"`
}

type Transactions []Transaction

type Amount struct {
	value    float64 `json:"value,omitempty"`
	currency string  `json:"currency,omitempty"`
}
