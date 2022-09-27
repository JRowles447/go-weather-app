package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Conf for app that includes API key for OpenWeather
type Conf struct {
	WeatherApiKey string
	Router        *gin.Engine
}

func ScaffoldApp() *Conf {
	conf := ParseEnv()
	conf.InitializeRouter()

	return conf
}

// parse environment variables needed for application
func ParseEnv() *Conf {
	conf := Conf{}
	conf.WeatherApiKey = os.Getenv("WEATHER_API_KEY")

	return &conf
}

func (conf *Conf) InitializeRouter() *gin.Engine {
	router := gin.Default()
	conf.Router = router

	router.GET("/zip", conf.GetCoordinatesHandler)
	router.GET("/weather", conf.QueryWeatherHandler)

	return router
}

// TODO is there a good way to remove the conf receiver?
func (conf *Conf) GetCoordinatesHandler(c *gin.Context) {
	var zip CoordinateRequest

	if err := c.ShouldBindJSON(&zip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad zip provided.",
			"error":   err.Error(),
		})
	}

	coordinates := conf.ConvertZipToCoordinates(zip.Zip)
	c.IndentedJSON(http.StatusOK, coordinates)
}

func (conf *Conf) QueryWeatherHandler(c *gin.Context) {
	var coordinates CoordinateResponse

	if err := c.ShouldBindJSON(&coordinates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad coordinates provided.",
			"error":   err.Error(),
		})
	}

	result := conf.QueryWeather(coordinates.Longitude, coordinates.Latitude)
	c.IndentedJSON(http.StatusOK, result)
}
