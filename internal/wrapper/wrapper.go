// This module contains the wrapper and data extraction of API calls
package wrapper

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"leoarayav/w/internal/exceptions"

	"github.com/joho/godotenv"
)

// Wrapp is the interface for the wrapper
//
// Methods:
//
//	Wrapper() error
//	LoadEnviroment()
type Wrapp interface {
	Wrapper() error
}

// This method is used to get the API endpoint
//
// Parameters:
//   - This method does not take any parameters
//
// Returns:
//   - This method returns the API endpoint
func Wrapper() error {
	err := godotenv.Load()
	if err != nil {
		return exceptions.Except(exceptions.W_API_ERROR, err.Error())
	}
	endpoint := os.Getenv("API_ENDPOINT")
	response, err := http.Get(endpoint)
	if err != nil {
		return exceptions.Except(exceptions.W_API_ERROR, err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return exceptions.Except(exceptions.W_API_ERROR, "API call failed")
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return exceptions.Except(exceptions.W_API_ERROR, err.Error())
	}
	fmt.Println(string(body))
	return nil
}
