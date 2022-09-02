package main

import (
	"fmt"

	"github.com/jrowles447/go-weather-app/app"
)

func main() {
	fmt.Println("Welcome to Go-Weather-App!")
	fmt.Printf("\tI can provide you with weather information based on your zip code\n\n")
	fmt.Printf("Please provide a zip code!\n")
	fmt.Printf("> ")

	var zip string

	// take user input for zip code
	fmt.Scanln(&zip)

	// validate that zip is 5 digits
	if validated := app.ValidateZip(zip); !validated {
		fmt.Printf("Zip code must be 5 digits, you provided: '%s'\n", zip)
	}

	fmt.Printf("You provided: '%s'", zip)
}
