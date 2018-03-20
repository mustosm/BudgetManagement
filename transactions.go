package main

import (
	"time"
)

// Transaction structur
type Transaction struct {
	TransactionID   string    `json:"transactionId,omitempty"`
	Description     string    `json:"description,omitempty"`
	Amount          Amount    `json:"amount,omitempty"`
	TransactionType string    `json:"transactionType,omitempty"`
	PaymentMeans    string    `json:"paymentMeans,omitempty"`
	CategoryID      string    `json:"categoryId,omitempty"`
	AccountID       string    `json:"accountId,omitempty"`
	Month           string    `json:"month,omitempty"`
	Year            string    `json:"year,omitempty"`
	CreationDate    time.Time `json:"creationDate,omitempty"`
	UpdateDate      time.Time `json:"updateDate,omitempty"`
}

// Transactions structur
type Transactions []Transaction

// Amount structur
type Amount struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}
