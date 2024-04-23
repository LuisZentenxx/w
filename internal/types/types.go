// This module contains the types used in the internal core of W
package types

type Weather struct {
	// The location of the weather
	Location Location `json:"location"`

	// The current weather
	Current Current `json:"current"`
}

type Location struct {
	// The name of the location
	Name string `json:"name"`

	// Country of the location
	Country string `json:"country"`

	// Region of the location
	Region string `json:"region"`

	// Latitude of the location
	Lat float64 `json:"lat"`

	// Longitude of the location
	Lon float64 `json:"lon"`

	// Timezone of the location
	Timezone string `json:"timezone_id"`

	// Local time of the location
	Localtime string `json:"localtime"`
}

type Current struct {
	// The temperature of the location (in Celsius obviously)
	Temperature float64 `json:"temperature"`

	// TODO: Add more fields respected to the current weather showed in the REQUEST
}
