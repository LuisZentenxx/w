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

	// The wind direction of the location (North, South, West, East)
	WindDir string `json:"wind_dir"`

	// The wind speed of the location (Km/h)
	WindSpeed float64 `json:"wind_speed"`

	//The wind degree
	WindDegree string `json:"wind_degree"`

	// Percent of humidity present in the air
	Humidity float64 `json:"humidity"`

	// The precipitation level (millimeters)
	Precip float64 `json:"precip"`

	// Ultraviolet index
	UvIndex int `json:"uv_index"`

	// The air pressure (millibar)
	Pressure float64 `json:"pressure"`

	// The cloud cover level in percentage
	Cloudcover float64 `json:"cloudcover"`

	// The "Feels Like" temperature
	FeelsLike float64 `json:"feelslike"`

	// The visibility level (kilometers)
	Visibility float64 `json:"visibility"`

	// TODO: Add more fields respected to the current weather showed in the REQUEST
}
