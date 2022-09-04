# go-weather-app
This app prompts a user for a zip code and provides them with current weather information based on their location. 

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

## TODO
* Check that provided zip is valid. 
* Identify weather API 
* Implement query 
* Implement parsing for results for user
* Add functionality to identify location of user based on IP. 