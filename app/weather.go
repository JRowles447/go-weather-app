package app

import (
	"regexp"
)

// RetrieveZip retrieves a zip from the user
func RetrieveZip() {

}

// QueryWeather queries the XXX for weather based on zip
func QueryWeather() {

}

// ValidateZip ensures the zip code has a valid 5 digit format
func ValidateZip(zip string) bool {
	regex := regexp.MustCompile(`^\d{5}$`)

	return regex.Match([]byte(zip))
}
