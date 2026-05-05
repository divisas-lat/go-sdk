package models

import "fmt"

type DivisasError struct {
	StatusCode int
	Message    string
}

func (e *DivisasError) Error() string {
	return fmt.Sprintf("divisas API error: %d - %s", e.StatusCode, e.Message)
}

// Ensure DivisasError implements the error interface
var _ error = (*DivisasError)(nil)
