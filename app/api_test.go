package app

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestParseEnv Tests that Conf is constructed correctly based on env var
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

// TestScaffoldApp Checks that environment variables parsed correctly and server started
func TestScaffoldApp(t *testing.T) {
	os.Setenv("WEATHER_API_KEY", "my-weather-api-key")
	defer os.Unsetenv("WEATHER_API_KEY")

	actualConf := ScaffoldApp()

	// check router is defined and set
	if actualConf.Router == nil {
		t.Errorf("Expected that router set on conf, but router not set")
	}

}
