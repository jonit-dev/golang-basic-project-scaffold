package validation

import "fmt"

// StringValidator provides string validation services.
type StringValidator struct {
	// Configuration or state here
}

// New creates a new instance of StringValidator.
func New() *StringValidator {
	return &StringValidator{}
}

// Validate checks if the given string meets the application-specific criteria.
func (v *StringValidator) Validate(s string) bool {
	// Implement validation logic specific to your application

	fmt.Printf("Length of the string is %v\n", len(s))

	return len(s) > 5
}
