package main

import (
	"encoding/json"
	"log"
)

type Sandwich struct {
	BreadType  string
	Fillings   []string
	Vegetables []string
	Extras     []string
}

type SandwichBuilder struct {
	sandwich Sandwich
}

func (sb *SandwichBuilder) SetBreadType(breadType string) *SandwichBuilder {
	sb.sandwich.BreadType = breadType
	return sb
}

func (sb *SandwichBuilder) SetFillings(fillings []string) *SandwichBuilder {
	sb.sandwich.Fillings = fillings
	return sb
}

func (sb *SandwichBuilder) SetVegetables(vegetables []string) *SandwichBuilder {
	sb.sandwich.Vegetables = vegetables
	return sb
}

func (sb *SandwichBuilder) SetExtras(extras []string) *SandwichBuilder {
	sb.sandwich.Extras = extras
	return sb
}

func (sb *SandwichBuilder) Build() Sandwich {
	return sb.sandwich
}

func main() {
	sandwichBuilder := SandwichBuilder{}
	sandwich := sandwichBuilder.
				SetBreadType("atta").
				SetFillings([]string{"cheese", "chutney"}).
				SetVegetables([]string{"tomato", "cucumber", "onion"}).
				SetExtras([]string{"sausce", "chesse"}).
				Build()

	sandwichBytes, err := json.Marshal(sandwich)
	if err != nil {
		log.Println("Error: " + err.Error())
		log.Fatal("exiting due to error...")
	}
	
	log.Println("Custom Sandwich: ", string(sandwichBytes))
	log.Fatal("exiting after success...")
}