package app

import (
	"regexp"
)

// RetrieveZip retrieves a zip from the user
func RetrieveZip() {

}

// ConvertZipToCoordinates calls the OpenWeather API to retreive latitude + longitude
// corresponding to a zip.
func ConvertZipToCoordinates(zip string) (float64, float64) {
	// create request to send to OpenWeather API
	// http://api.openweathermap.org/geo/1.0/zip?zip=<ZIP>,US&appid=<APPID>

	return 0.0, 0.0
}

// QueryWeather queries the OpenWeather API for weather based on latitude + longitude
func QueryWeather() {

}

// ValidateZip ensures the zip code has a valid 5 digit format
func ValidateZip(zip string) bool {
	regex := regexp.MustCompile(`^\d{5}$`)

	return regex.Match([]byte(zip))
}
