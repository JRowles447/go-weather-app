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

	PrintWelcomeMessage()

	fmt.Println("Welcome to Go-Weather-App!")
	fmt.Printf("\tI can provide you with weather information based on your zip code\n\n")

	// scaffold application, parse env, setup router
	conf := app.ScaffoldApp()

	// start the server
	conf.Router.Run(":8080")
}

func PrintWelcomeMessage() {
	msg := `
	█     █░▓█████ ▄▄▄     ▄▄▄█████▓ ██░ ██ ▓█████  ██▀███  
	▓█░ █ ░█░▓█   ▀▒████▄   ▓  ██▒ ▓▒▓██░ ██▒▓█   ▀ ▓██ ▒ ██▒
	▒█░ █ ░█ ▒███  ▒██  ▀█▄ ▒ ▓██░ ▒░▒██▀▀██░▒███   ▓██ ░▄█ ▒
	░█░ █ ░█ ▒▓█  ▄░██▄▄▄▄██░ ▓██▓ ░ ░▓█ ░██ ▒▓█  ▄ ▒██▀▀█▄  
	░░██▒██▓ ░▒████▒▓█   ▓██▒ ▒██▒ ░ ░▓█▒░██▓░▒████▒░██▓ ▒██▒
	░ ▓░▒ ▒  ░░ ▒░ ░▒▒   ▓▒█░ ▒ ░░    ▒ ░░▒░▒░░ ▒░ ░░ ▒▓ ░▒▓░
	  ▒ ░ ░   ░ ░  ░ ▒   ▒▒ ░   ░     ▒ ░▒░ ░ ░ ░  ░  ░▒ ░ ▒░
	  ░   ░     ░    ░   ▒    ░       ░  ░░ ░   ░     ░░   ░ 
		░       ░  ░     ░  ░         ░  ░  ░   ░  ░   ░     
															 
	`

	fmt.Println(msg)
}
