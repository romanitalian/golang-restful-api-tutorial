package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"bytes"
)

const mugQuantity = 50

func GetInfo(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get infos about beers")
	initHeaders(writer)
	json.NewEncoder(writer).Encode(beers)
}

func GetBeerInfo(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	//Converts the id parameter from a string to an int
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err == nil {
		log.Println("Get info about beer id #", id)

		//Retrieves the infos about the beer
		beer := FindBeerByID(id)
		json.NewEncoder(writer).Encode(beer)
	} else {
		log.Fatal(err.Error())
	}
}

func OrderBeer(writer http.ResponseWriter, request *http.Request) {
	log.Println("Order a beer")
	initHeaders(writer)
	var order Order

	//Decodes the request and put the content of the body into the order
	_ = json.NewDecoder(request.Body).Decode(&order)

	//Retrieves the infos about the beer he wants to order
	beer := FindBeerByID(order.ID)

	numberOfBeerWanted := order.Quantity / mugQuantity
	//If the customer sends enough money
	//float32() converts a int into a float32
	if order.Money >= beer.Price * float32(numberOfBeerWanted) {
		mugs := serveBeer(&order, numberOfBeerWanted)

		json.NewEncoder(writer).Encode(mugs)
	} else {
		json.NewEncoder(writer).Encode("No enough money")
	}
}

func BreakMug(writer http.ResponseWriter, request *http.Request) {
	log.Println("A mug broke")
	initHeaders(writer)
	numberOfBrokenMug++
	json.NewEncoder(writer).Encode("Tonight, " + strconv.Itoa(numberOfBrokenMug) + " mug(s) broke")
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func serveBeer(order *Order, numberOfBeerWanted int) []Mug {
	var mugs []Mug

	//We search for the good barrel
	barrel, idx := FindBarrelFromBeerID(order.ID, barrels)
	for i := 0; i < numberOfBeerWanted; i++ {
		var mug Mug
		//When there is no more enough beer, we order a new barrel
		if (barrels[idx].Quantity - mugQuantity) <= 0 {
			orderBarrel(idx, barrel)
		}

		mug.Beer = barrel.Beer
		mug.Quantity = mugQuantity
		barrels[idx].Quantity -= mugQuantity
		mugs = append(mugs, mug)
	}


	log.Println("It left", barrels[idx].Quantity, "cl in the barrel of", barrel.Beer.Name)
	return mugs
}

func orderBarrel(idx int, barrel Barrel) {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(barrel)
	res, _ := http.Post("http://localhost:8000/brewery", "application/json", buffer)
	var newBarrel Barrel
	json.NewDecoder(res.Body).Decode(&newBarrel)
	barrels[idx].Quantity += newBarrel.Quantity

	log.Println("The barrel of ", barrel.Beer.Name, " has been refilled, it has now", barrels[idx].Quantity, "cl")
}

