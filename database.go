package main

import "time"

var beers []Beer
var barrels []Barrel
var breweryBarrels []Barrel
var numberOfBrokenMug int

type Beer struct {
	ID int   `json:"id"`
	Name string   `json:"name,omitempty"`
	Price  float32   `json:"price,omitempty"`
	PercentProof float32 `json:"percentProof,omitempty"`
	IPA bool `json:"ipa"`
}

type Barrel struct {
	Beer *Beer `json:"beer"`
	Quantity int `json:"quantity"`
	DateOfManufacture time.Time `json:"dateOfManufacture"`
}

type Mug struct {
	Beer *Beer `json:"beer"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID int `json:"id"`
	Money float32 `json:"money"`
	Quantity int `json:"quantity"`
}

// Initializing datas
func initDatas() {
	//Beers
	gandalf := Beer{ID: 0, Name: "Gandalf", Price: 5, PercentProof:8, IPA: false}
	aragorn := Beer{ID: 1, Name: "Aragorn", Price: 5.5, PercentProof:7.5, IPA: true}
	sauron := Beer{ID: 2, Name: "Sauron", Price: 7, PercentProof: 11, IPA: false}

	beers = append(beers, gandalf, aragorn, sauron)

	//Barrels
	barrels = append(barrels,
		Barrel{&gandalf, 1000, time.Now()},
		Barrel{&aragorn, 8000, time.Now()},
		Barrel{&sauron, 0, time.Now()})

	//Brewery Barrels
	breweryBarrels = append(breweryBarrels,
		Barrel{&gandalf, 1000, time.Now()},
		Barrel{&gandalf, 1000, time.Now()},
		Barrel{&gandalf, 1000, time.Now()},
		Barrel{&aragorn, 5000, time.Now()},
		Barrel{&aragorn, 5000, time.Now()},
		Barrel{&sauron, 3000, time.Now()})
}

func FindBeerByID(ID int) Beer {
	var result Beer
	for _, beer := range beers {
		if beer.ID == ID {
			result = beer
			break
		}
	}
	return result
}

func FindBarrelFromBeerID(ID int, barrels []Barrel) (Barrel, int) {
	var result Barrel
	var idx = -1
	for i, barrel := range barrels {
		if barrel.Beer.ID == ID {
			result = barrel
			idx = i
			break
		}
	}
	return result, idx
}

func FindBarrelFromBreweryByBeer(beer *Beer) (Barrel, int) {
	return FindBarrelFromBeerID(beer.ID, breweryBarrels)
}