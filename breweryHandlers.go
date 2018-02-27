package main

import (
	"net/http"
	"log"
	"encoding/json"
	"time"
)

func OrderBarrels(writer http.ResponseWriter, request *http.Request) {
	log.Println("Order new Barrel")
	initHeaders(writer)

	var requestedBarrel Barrel
	_ = json.NewDecoder(request.Body).Decode(&requestedBarrel)

	//Try to find in stock the barrel.
	barrel, idx := FindBarrelFromBreweryByBeer(requestedBarrel.Beer)

	//If idx is inferior than 0, we need to produce new barrels
	if idx < 0 {
		//Initialize a client
		client := &http.Client{}
		//Prepare a PUT Request to http://localhost:8000/brewery with no body
		request, _ := http.NewRequest(http.MethodPut, "http://localhost:8000/brewery", nil)
		//Send the request
		client.Do(request)
		//"Reload" the barrel
		barrel, idx = FindBarrelFromBreweryByBeer(requestedBarrel.Beer)
	}

	requestedBarrel = barrel
	//Removes the barrel from the stock
	breweryBarrels = append(breweryBarrels[:idx], breweryBarrels[idx+1:]...)

	json.NewEncoder(writer).Encode(requestedBarrel)
}

func ProduceBarrels(writer http.ResponseWriter, request *http.Request) {
	log.Println("Producing new Barrels")

	breweryBarrels = append(breweryBarrels,
		Barrel{&beers[0], 1000, time.Now()},
		Barrel{&beers[1], 5000, time.Now()},
		Barrel{&beers[2], 3000, time.Now()})
}
