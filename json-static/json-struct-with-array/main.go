package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

// Stars struct to hold a slice of Star structs
type Stars struct {
	StarList []Star `json:"stars"`
}

// Star struct for each item into the array
type Star struct {
	Name          string  `json:"name"`
	Distance      float64 `json:"distance"`
	Constellation string  `json:"constellation"`
}

func main() {
	fmt.Println("Marshalling and unmarshalling the brightest stars to/from a JSON array")

	// Create slice of "Star" structs
	var structStars = Stars{
		StarList: []Star{
			Star{Name: "Sun", Distance: 0.000015813, Constellation: ""},
			Star{Name: "Sirius", Distance: 8.6, Constellation: "Canis Major"},
			Star{Name: "Canopus", Distance: 310, Constellation: "Carina"},
			Star{Name: "RigilKentaurus", Distance: 4.4, Constellation: "Centaurus"},
			Star{Name: "Toliman", Distance: 4.4, Constellation: "Centaurus"},
			Star{Name: "Arcturus", Distance: 37, Constellation: "Boötes"},
		},
	}

	fmt.Printf("\n* Initial \"Star\" struct slice:\n%+v\n", structStars)
	marshalled := encode(structStars) // Marshalling
	fmt.Printf("\n* \"Stars\" struct marshalled to JSON string:\n%s\n", marshalled)
	unmarshalled := decode(marshalled) // Unmarshalling
	fmt.Printf("\n* JSON string unmarshalled to \"Stars\" struct:\n%+v\n", unmarshalled)
	fmt.Printf("\n* Iterating over newly created \"Stars\" struct: \n")
	for _, star := range unmarshalled.StarList { // Iterating slice of structs
		fmt.Printf("\n")
		fmt.Printf("\tName: %s\n", star.Name)
		fmt.Printf("\tDistance: %.2f ly\n", math.Round(star.Distance*100)/100)
		fmt.Printf("\tConstellation: %s\n", star.Constellation)
	}
	fmt.Printf("\n")
}

// Gets a slice of "Star" structs and returns it's JSON string representation
func encode(structStars Stars) string {
	// Marshalling + error control
	jsonStars, err := json.Marshal(structStars)
	if err != nil {
		log.Fatalln("Cannot encode to JSON ", err)
	}

	return string(jsonStars)
}

// Gets a JSON string and returns it's "Star" struct representation
func decode(jsonInput string) Stars {
	// Converting JSON string to a slice of type byte,
	// required by the function json.Unmarshal
	jsonStars := []byte(jsonInput)

	// Creating an array of struct Start to store Unmarshalled items
	var structStars Stars

	// Unmarshalling + error control
	if err := json.Unmarshal(jsonStars, &structStars); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s\n", err)
	}

	return (structStars)
}
