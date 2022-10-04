# go-weather-app [![build](https://github.com/JRowles447/go-weather-app/actions/workflows/package.yaml/badge.svg?branch=main)](https://github.com/JRowles447/go-weather-app/actions/workflows/package.yaml)

This app runs a gin server with endpoints to query weather for a provided zip code. Under the hood it leverages the OpenWeather API. For more information about the OpenWeather API, please refer to the following: [OpenWeather](https://openweathermap.org/).

## To Run
In order to run the application, clone the repository and run the following from the project root: 
```
go run . 
```

### Or, run it in Docker! 

Build the image:
```
docker build --tag go-weather-app .

```

Run the image: 
```
docker run go-weather-app:latest
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

## Building a Docker Image
In order to build a Docker image for the project, run the following from the project root directory: 
```
docker build --tag go-weather-app .
```

## Sample Output
The following is sample output from running the application: 
```
Welcome to Go-Weather-App!
	I can provide you with weather information based on your zip code

Please provide a zip code!
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /zip                      --> github.com/jrowles447/go-weather-app/app.(*Conf).GetCoordinatesHandler-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
```

If your application is already leveraging port 8080, you may need to update the configuration to leverage a different port number. 

## TODO
* Add Swagger Spec
* Add functionality to identify location of user based on IP. 