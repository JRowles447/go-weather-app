package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Go-Weather-App!")
	fmt.Printf("\tI can provide you with weather information based on your zip code\n\n")
	fmt.Printf("Please provide a zip code!\n> ")

	var zip string

	fmt.Scanln(&zip)
	fmt.Printf("You provided: '%s'", zip)
}
