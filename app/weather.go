package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// ConvertZipToCoordinates calls the OpenWeather API to retreive latitude + longitude
// corresponding to a zip.
// More information about OpenWeather zip converter API: https://openweathermap.org/api/geocoding-api#direct_name_how
func (conf *Conf) ConvertZipToCoordinates(zip string) CoordinateResponse {
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

	res := CoordinateResponse{
		Latitude:  openWeatherResp.Latitude,
		Longitude: openWeatherResp.Longitude,
	}

	return res
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
		fmt.Printf("Error reading response from OpenWeatherAPI current weather call")
	}

	fmt.Printf("result: '%+v'", string(body))
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
