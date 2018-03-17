package budget

import (
    "encoding/json"
    "net/http"
)

func GetTransationById(w http.ResponseWriter, r *http.Request) {
    transaction := Transaction{

    }
    if err := json.NewEncoder(w).Encode(transactions); err != nil {
        panic(err)
    }
}

func GetTransations(w http.ResponseWriter, r *http.Request) {
    transactions := Transaction{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    if err := json.NewEncoder(w).Encode(transactions); err != nil {
        panic(err)
    }
}