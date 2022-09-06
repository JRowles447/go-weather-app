package main

// type OpenWeatherCoordinateRequest struct {

// }

type OpenWeatherCoordinateResponse struct {
	Zip       string  `json:"zip"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Country   string  `json:"country"`
}
