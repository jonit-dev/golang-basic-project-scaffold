package main

import (
	"fmt"
	"jonit-dev/golang-basic-project-scaffold/internal/validation"
)

func main() {

	fmt.Println("Hello World!")

	// Create a new instance of StringValidator
	v := validation.New()

	// Validate a string
	if v.Validate("Hello World from StringValidator!") {
		fmt.Println("Valid string")
	} else {
		fmt.Println("Invalid string")
	}

}
