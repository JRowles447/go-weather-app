package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Conf struct {
	WeatherApiKey string
}

// parse environment variables needed for application
func ParseEnv() *Conf {
	conf := Conf{}
	conf.WeatherApiKey = os.Getenv("WEATHER_API_KEY")

	fmt.Printf("WEATHER_API: %s\n", conf.WeatherApiKey)
	return &conf
}

// RetrieveZip retrieves a zip from the user
func RetrieveZip() {

}

// ConvertZipToCoordinates calls the OpenWeather API to retreive latitude + longitude
// corresponding to a zip.
func (conf *Conf) ConvertZipToCoordinates(zip string) (float64, float64) {
	// create request to send to OpenWeather API
	// http://api.openweathermap.org/geo/1.0/zip?zip=<ZIP>,US&appid=<APPID>
	client := &http.Client{}

	endpoint := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/zip?zip=%s,US&appid=%s", zip, conf.WeatherApiKey)
	req, err := http.NewRequest("GET", endpoint, nil)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error sending request to endpoint")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response from OpenWeatherAPI zip converter")
	}

	var openWeatherResp OpenWeatherCoordinateResponse
	if err := json.Unmarshal(body, &openWeatherResp); err != nil {
		fmt.Println("Issue unmarshalling response from OpenWeatherAPI zip converter")
	}

	fmt.Printf("Zip:%s, Longitude: %f, Latitude: %f\n", zip, openWeatherResp.Longitude, openWeatherResp.Latitude)
	return openWeatherResp.Longitude, openWeatherResp.Latitude
}

// QueryWeather queries the OpenWeather API for weather based on latitude + longitude
func QueryWeather(long float64, lat float64) int {

	return 0
}

// ValidateZip ensures the zip code has a valid 5 digit format
func ValidateZip(zip string) bool {
	regex := regexp.MustCompile(`^\d{5}$`)

	return regex.Match([]byte(zip))
}
