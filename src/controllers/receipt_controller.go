package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"receipt-processor-challenge/src/entities"
	"receipt-processor-challenge/src/errors"
	"receipt-processor-challenge/src/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/exp/maps"
)

var Receipts = make(map[string]entities.Receipt)

func generateUUID() string {
	return uuid.New().String()
}

// POST - create new receipt, with random ID - saved in memory
func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	log.Printf("Creating a new receipt..")

	w.Header().Set("Content-Type", "application/json")
	var receipt entities.Receipt
	_ = json.NewDecoder(r.Body).Decode(&receipt)
	receipt.ID = generateUUID()
	Receipts[receipt.ID] = receipt
	log.Printf("Created receipt " + receipt.ID)
	json.NewEncoder(w).Encode(receipt)
}

// GET - get all receipts
func GetAllReceipts(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting all receipts saved in memory")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(maps.Values(Receipts))
}

// GET - single receipt
func GetReceipt(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("Getting receipt with ID: " + params["id"])
	receipt, exists := Receipts[params["id"]]
	if exists {
		json.NewEncoder(w).Encode(receipt)
		return
	}
	errors.ErrorHandler(w, r, 404)
}

// GET - get points for the receipt
func CalculatePointsForReceipt(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("Calculating points for receipt ID: " + params["id"])
	receipt, exists := Receipts[params["id"]]
	if exists {
		json.NewEncoder(w).Encode(utils.Calculate(receipt))
		return
	}
	errors.ErrorHandler(w, r, 404)
}
