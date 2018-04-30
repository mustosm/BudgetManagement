package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TransactionsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "transactions"
)

func (trs *TransactionsDAO) Connect() {
	session, err := mgo.Dial(trs.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(trs.Database)
}

func (trs *TransactionsDAO) FindAll() (Transactions, error) {
	var transactions Transactions
	err := db.C(COLLECTION).Find(bson.M{}).All(&transactions)
	return transactions, err
}

func (m *TransactionsDAO) FindById(id string) (Transaction, error) {
	var transaction Transaction
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&transaction)
	return transaction, err
}

func (m *TransactionsDAO) Insert(transaction Transaction) error {
	err := db.C(COLLECTION).Insert(&transaction)
	return err
}

func (m *TransactionsDAO) Delete(transaction Transaction) error {
	err := db.C(COLLECTION).RemoveId(transaction.TransactionID)
	return err
}

func (m *TransactionsDAO) Update(transaction Transaction, transactionId string) error {
	ID := bson.ObjectIdHex(transactionId)
	err := db.C(COLLECTION).UpdateId(ID, &transaction)
	return err
}

func (trs *TransactionsDAO) HealthCheck() ([]string, error) {
	collections, err := db.CollectionNames()
	return collections, err
}