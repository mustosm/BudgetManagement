package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Transaction structure
type Transaction struct {
	TransactionID bson.ObjectId `bson:"_id" json:"transactionId,omitempty"`
	Description   string        `bson:"description" json:"description,omitempty"`
	Amount        *Amount       `bson:"amount" json:"amount,omitempty"`
	Type          string        `bson:"type" json:"Type,omitempty"`
	PaymentMeans  string        `bson:"paymentMeans" json:"paymentMeans,omitempty"`
	CategoryID    string        `bson:"categoryId" json:"categoryId,omitempty"`
	AccountID     string        `bson:"accountId" json:"accountId,omitempty"`
	Month         string        `bson:"month" json:"month,omitempty"`
	Year          string        `bson:"year" json:"year,omitempty"`
	CreationDate  time.Time     `bson:"creationDate" json:"creationDate,omitempty"`
	UpdateDate    time.Time     `bson:"updateDate" json:"updateDate,omitempty"`
}

// Transactions structure
type Transactions []Transaction

// Amount structure
type Amount struct {
	Value    float64 `bson:"value" json:"value,omitempty"`
	Currency string  `bson:"currency" json:"currency,omitempty"`
}
