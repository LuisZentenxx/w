// This module contains the wrapper and data extraction of API calls
package wrapper

import (
	"fmt"
	"net/http"
	"os"

	"leoarayav/w/internal/exceptions"

	"github.com/joho/godotenv"
)

var (
	// API_ENDPOINT is the endpoint of the API
	API_ENDPOINT = os.Getenv("API_ENDPOINT")
)

// This method is used to load the environment variables
//
// Parameters:
//   - This method does not take any parameters
//
// Returns:
//   - This method does not return any values
func LoadEnviroment() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("> Error loading .env file")
	}
	endpoint := os.Getenv("API_ENDPOINT")
	fmt.Println("> API Endpoint was loaded successfully", endpoint)
}

// This method is used to get the API endpoint
//
// Parameters:
//   - This method does not take any parameters
//
// Returns:
//   - This method returns the API endpoint
func Wrapper() exceptions.ErrorProps {
	res, err := http.Get(API_ENDPOINT)
	if err != nil {
		return exceptions.Except(exceptions.W_API_ERROR, "Error getting the API endpoint")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return exceptions.Except(exceptions.W_API_ERROR, "Error getting the API endpoint")
	}
	return exceptions.Except(exceptions.W_API_ERROR, "Error getting the API endpoint")
}
