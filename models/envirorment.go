package models

// Environment (Live or Test)
type Environment = string

const (
	// NotSelected -
	NotSelected Environment = "none"

	// Live environment
	Live Environment = "live"

	// Test environment
	Test Environment = "test"
)
