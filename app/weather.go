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

	return &conf
}

// RetrieveZip retrieves a zip from the user
func RetrieveZip() {

}

// ConvertZipToCoordinates calls the OpenWeather API to retreive latitude + longitude
// corresponding to a zip.
// More information about OpenWeather zip converter API: https://openweathermap.org/api/geocoding-api#direct_name_how
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
// For more information about the OpenWeather current weather API, refer to the following: https://openweathermap.org/current
func (conf *Conf) QueryWeather(long float64, lat float64) OpenWeatherCurrentWeatherResult {
	client := &http.Client{}

	endpoint := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=imperial&appid=%s", lat, long, conf.WeatherApiKey)
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

	var openWeatherResp OpenWeatherCurrentWeatherResult
	if err := json.Unmarshal(body, &openWeatherResp); err != nil {
		fmt.Println("Issue unmarshalling response from OpenWeatherAPI current weather call")
	}
	fmt.Printf("City: %s, Longitude: %f, Latitude: %f, Temperature: %f\n", openWeatherResp.Name, openWeatherResp.Coordinates.Longitude, openWeatherResp.Coordinates.Latitude, openWeatherResp.Main.Temp)

	return openWeatherResp
}

// ValidateZip ensures the zip code has a valid 5 digit format
func ValidateZip(zip string) bool {
	regex := regexp.MustCompile(`^\d{5}$`)

	return regex.Match([]byte(zip))
}

// FormatWeatherResultString constructs a string to return to the user with pertinent information
// from the response of the OpenWeather current weather call.
func FormatWeatherResultString(currentWeather OpenWeatherCurrentWeatherResult) string {
	ret := `
	City: %s
	Weather Description: %s
	Temperature: %f
	Feels Like: %f
	Humidity: %d
	Wind Speed: %f
	`
	res := fmt.Sprintf(ret, currentWeather.Name, currentWeather.Weather[0].Description, currentWeather.Main.Temp, currentWeather.Main.FeelsLike, currentWeather.Main.Humidity, currentWeather.Wind.Speed)
	return res
}
