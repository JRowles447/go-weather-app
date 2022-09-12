package app

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseEnv(t *testing.T) {
	os.Setenv("WEATHER_API_KEY", "my-weather-api-key")
	defer os.Unsetenv("WEATHER_API_KEY")

	actual := *ParseEnv()

	expected := Conf{
		WeatherApiKey: "my-weather-api-key",
	}

	if !cmp.Equal(actual, expected) {
		t.Errorf("got %+v, want %+v", actual, expected)
	}
}

func TestValidateZip_Valid(t *testing.T) {
	var tests = []struct {
		zip  string
		want bool
	}{
		{"80209", true},
		{"98109", true},
		{"02210", true},
		{"02144", true},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.zip)

		t.Run(testName, func(t *testing.T) {
			res := ValidateZip(tt.zip)
			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}
}

// Contains mix of valid and invalid digit strings
func TestValidateZip_InvalidDigits(t *testing.T) {
	var tests = []struct {
		zip  string
		want bool
	}{
		{"12", false},
		{"557852", false},
		{"90000", true},
		{"0", false},
		{"69621475", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.zip)

		t.Run(testName, func(t *testing.T) {
			res := ValidateZip(tt.zip)
			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}
}

// Contains a mix of valid zips and invalid character strings
func TestValidateZip_InvalidCharacters(t *testing.T) {
	var tests = []struct {
		zip  string
		want bool
	}{
		{"foo", false},
		{"another", false},
		{"", false},
		{"bar", false},
		{"12345", true},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.zip)

		t.Run(testName, func(t *testing.T) {
			res := ValidateZip(tt.zip)
			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}
}
