package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Go-Weather-App!")
	fmt.Printf("\tI can provide you with weather information based on your zip code\n\n")
	fmt.Printf("Please provide a zip code!\n")
	fmt.Printf("> ")

	var zip string
	var parsedZip int64

	// take user input for zip code
	fmt.Scanln(&zip)

	// parse zip code value to int
	parsedZip, err := strconv.ParseInt(zip, 10, 64)
	if err != nil { // if there is a parsing error, exit
		fmt.Printf("You did not provide a valid zip code, please try again\n")
		os.Exit(-1)
	}

	fmt.Printf("You provided: '%d'", parsedZip)
}
