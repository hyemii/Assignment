package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Inventory struct {
	Vin string `json:"vin"`
	Model string `json:"model"`
	Make string `json:"make"`
	Year int `json:"year"`
	MSRP int `json:"msrp"`
	Status string `json:"status"`
	Booked string `json:"booked"`
	Listed string `json:"listed"`
}

var inventories []Inventory

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
	(*w).Header().Set("Content-Type", "application/json")
}

func getInventories(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	json.NewEncoder(w).Encode(inventories)
}

func createInventory(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	var inventory Inventory
	_ = json.NewDecoder(r.Body).Decode(&inventory)

	for _ , Inventory := range inventories {
		err := Inventory.Vin == inventory.Vin
		if err {
			fmt.Println("Vin confirm error")
			return
		}
	}

	inventories = append(inventories, inventory)
	json.NewEncoder(w).Encode(&inventory)
}

func deleteInventories(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	params := mux.Vars(r)
	delvins := params["delvins"]
	var arrDelVins []string
	var converFlag = strings.ContainsAny(delvins, ",")

	if !converFlag {
		arrDelVins = append(arrDelVins, delvins)
	}

	if converFlag {
		arrDelVins = strings.Split(delvins, ",")
	}

	for index, inventory := range inventories {
		for _ , delvin := range arrDelVins {
			if inventory.Vin == delvin {
				inventories = append(inventories[:index], inventories[index+1:]...)
				break
			}
		}
	}
	json.NewEncoder(w).Encode(inventories)
}

func main() {
	router := mux.NewRouter()

	inventories = append(inventories, Inventory{Vin: "TFBAXXMAWAFS71274", Model: "Focus", Make: "Ford", Year: 2016, MSRP: 130000, Status: "ordered", Booked: "N", Listed: "Y"})
	inventories = append(inventories, Inventory{Vin: "MNBUMF050FW496402", Model: "320i", Make: "BMW", Year: 2014, MSRP: 147000, Status: "ordered", Booked: "Y", Listed: "N"})
	inventories = append(inventories, Inventory{Vin: "4JDBLMF080FW468775", Model: "Camry", Make: "Toyota", Year: 2015, MSRP: 120000, Status: "in stock", Booked: "Y", Listed: "N"})

	router.HandleFunc("/inventories", getInventories).Methods("GET", "OPTIONS")
	router.HandleFunc("/inventories", createInventory).Methods("POST", "OPTIONS")
	router.HandleFunc("/inventories/{delvins}", deleteInventories).Methods("DELETE", "OPTIONS")
	http.ListenAndServe(":8090", router)
}