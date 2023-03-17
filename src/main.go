package main

import (
	"fmt"
	"log"
	"net/http"
	"receipt-processor-challenge/src/controllers"

	// TODO using a router that is not being maintained anymore, this needs to be researched and replaced if necessary
	"github.com/gorilla/mux"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

func fetch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Fetch!")
}

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)
	/**
	{
		"retailer": "Target", // 6
		"purchaseDate": "2022-01-02", // +6 =12
		"purchaseTime": "15:13", // +10 = 49
		"total": "1.25", // +25 = 31
		"items": [
			{"shortDescription": "Pepsi - 12oz  ", "price": "1.25"} //+0 since its 0.2
		]
	}
	*/
	handleReceiptRequests(router)
}

func handleReceiptRequests(r *mux.Router) {

	http.HandleFunc("/home", fetch)
	r.HandleFunc("/receipts", controllers.GetAllReceipts).Methods("GET")
	r.HandleFunc("/receipts/{id}", controllers.GetReceipt).Methods("GET")
	r.HandleFunc("/receipts/{id}/points", controllers.CalculatePointsForReceipt).Methods("GET")
	r.HandleFunc("/receipts", controllers.CreateReceipt).Methods("POST")
	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
	fmt.Println("Starting server at port:", Port)
	log.Fatal(http.ListenAndServe(": "+Port, r))
}
