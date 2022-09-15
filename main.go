package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/jrowles447/go-weather-app/app"
)

func main() {
	// load local env file necessary to run application
	err := godotenv.Load("env.local")
	if err != nil {
		fmt.Printf("Issue parsing env file!\n")
	}

	// scaffold application, parse env, setup router
	conf := app.ScaffoldApp()

	// start the server
	conf.Router.Run(":8080")

	fmt.Println("Welcome to Go-Weather-App!")
	fmt.Printf("\tI can provide you with weather information based on your zip code\n\n")
	fmt.Printf("Please provide a zip code!\n")
	fmt.Printf("> ")

	// var zip string

	// // take user input for zip code
	// fmt.Scanln(&zip)

	// // validate that zip is 5 digits
	// if validated := app.ValidateZip(zip); !validated {
	// 	fmt.Printf("Zip code must be 5 digits, you provided: '%s'\n", zip)
	// } else {
	// 	fmt.Printf("You provided: '%s'\n", zip)

	// 	coordinates := conf.ConvertZipToCoordinates(zip)

	// 	currWeather := conf.QueryWeather(coordinates.Longitude, coordinates.Latitude)

	// 	// print current weather response formatted
	// 	fmt.Printf(app.FormatWeatherResultString(currWeather))
	// }
}
