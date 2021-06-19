package status

import (
	"errors"
	"strconv"
)

// Code - used for our custom error codes to specify for the forntend what exactly happened
type Code int

const (
	// Core Statuses (1000-1099)

	//EmptyBody - error when body is not supplied
	EmptyBody = 1000
	//IncorrectBodyFormat - error when decoded body is not in correct format
	IncorrectBodyFormat = 1001
	//ErrorMissingWord - error when word is not in correct format
	ErrorMissingWord = 1002
	//ErrorSynonymAllreadyEsists - error when user is trying to add already existing synonym
	ErrorSynonymAllreadyEsists = 1003
	//ErrorWordDoesNotExist - word does not exist
	ErrorWordDoesNotExist = 1004
)

/// ****************************************************
/// INFO
/// ****************************************************
var statusText = map[int]string{
	EmptyBody:                  "Empty request body",
	IncorrectBodyFormat:        "Incorrect body format",
	ErrorMissingWord:           "Word is not in correct format",
	ErrorSynonymAllreadyEsists: "The word alreeay exists as a synonym",
	ErrorWordDoesNotExist:      "The word does not exist",
}

var statusErrors = map[int]error{
	EmptyBody:           errors.New(statusText[EmptyBody]),
	IncorrectBodyFormat: errors.New(statusText[IncorrectBodyFormat]),
}

/// ************************************************
///     HELPER METHODS
/// ************************************************

// Error - Retrieves the error for the specific code
func Error(code int) error {
	return statusErrors[code]
}

//Text - Retrieces the message for the specific code
func Text(code int) string {
	return statusText[code]
}

// StringCode converts existing number into string
func StringCode(code int) string {
	return strconv.Itoa(code)
}

// Itoc convertsa an integer into `Code`
func Itoc(code int) Code {
	return Code(code)
}
