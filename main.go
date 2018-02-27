package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main(){
	fmt.Println("Welcome to the The Green Dragon !")
	router := mux.NewRouter()

	initDatas()
	buildBarRoutes(router)
	buildBreweryRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}

//Bar's routes
func buildBarRoutes(router *mux.Router) {
	prefix := "/bar"
	router.HandleFunc(prefix, GetInfo).Methods("GET")
	router.HandleFunc(prefix + "/{id}", GetBeerInfo).Methods("GET")
	router.HandleFunc(prefix, OrderBeer).Methods("POST")
	router.HandleFunc(prefix, BreakMug).Methods("DELETE")
}

//Brewery's routes
func buildBreweryRoutes(router *mux.Router) {
	prefix := "/brewery"
	router.HandleFunc(prefix, OrderBarrels).Methods("POST")
	router.HandleFunc(prefix, ProduceBarrels).Methods("PUT")
}


