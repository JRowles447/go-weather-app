# go-weather-app
This app prompts a user for a zip code and provides them with current weather information based on their location. Under the hood it leverages the OpenWeather API. For more information about the OpenWeather API, please refer to the following: [OpenWeather](https://openweathermap.org/).

## To Run
In order to run the application, clone the repository and run the following from the project root: 
```
go run . 
```

## To Test
To test the application, run the following: 
```
go test ./app
```

To generate and view a coverage report, run the following: 
```
go test -coverprofile=coverage.out ./app
go tool cover -html=coverage.out
```

## Sample Output
The following is sample output from running the application: 
```
Welcome to Go-Weather-App!
	I can provide you with weather information based on your zip code

Please provide a zip code!
> 98121
You provided: '98121'
Zip:98121, Longitude: -122.344700, Latitude: 47.615100
City: Seattle, Longitude: -122.344700, Latitude: 47.615100, Temperature: 75.700000

	City: Seattle
	Weather Description: smoke
	Temperature: 75.700000
	Feels Like: 75.510000
	Humidity: 54
	Wind Speed: 4.000000
```

If there is a problem with the provided zip code, the following will be returned: 
```
Welcome to Go-Weather-App!
	I can provide you with weather information based on your zip code

Please provide a zip code!
> badzip
Zip code must be 5 digits, you provided: 'badzip'
```
## TODO
* Add functionality to identify location of user based on IP. 