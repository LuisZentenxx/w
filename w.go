package main

import (
	"fmt"
	"leoarayav/w/internal/wrapper"
)

// This is the main package of the application
//
// Usage: w --help | --version | [country]
//
// Options:
//
//	--help     Show this screen.
//	--version  Show version.
//
// Arguments:
//
//	country    Country name where to get the weather information.
func main() {
	w := wrapper.Wrapper()
	if w != nil {
		fmt.Println(w)
	}
}
