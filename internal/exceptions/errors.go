// This module contains the exceptions and error handling resources
package exceptions

import "fmt"

// Error type definition
//
// Fields:
// 	Type    string
// 	Message string
type ErrorProps struct {
	Type    string
	Message string
}

// Existing error scenarios
var (
	// This error is thrown when the country is not found
	W_UNKNOWN_COUNTRY = "UNKNOWN_COUNTRY"

	// This error is thrown when the API call fails
	W_API_ERROR = "API_CALL_ERROR"

	// This error is thrown when an error occurs internally
	W_ERROR = "INTERNAL_ERROR"

	// This error is thrown when no data is found
	W_NO_DATA = "NO_DATA"

	// This error is thrown when reaching the maximum retries
	W_MAX_RETRIES = "MAX_RETRIES"
)

// This method is used to get the error message
//
// Parameters:
//   - This method does not take any parameters
//
// Returns:
//   - This method returns the error message
func (e *ErrorProps) Error() string {
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

// This method is used to create a new error
//
// Parameters:
//  - text: The error message
//
// Returns:
//   - This method returns the error properties
func Except(typo string, message string) error {
	return &ErrorProps{typo, message}
}
