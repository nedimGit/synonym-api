package viewmodels

import (
	"encoding/json"
	"log"
	"net/http"

	status "github.com/NedimUka/synonyms/viewmodels/statusCodes"
)

// AddWordRequest - model
type AddWordRequest struct {
	Word      Word  `json:"Word"`
	Timestamp int64 // setup on backend side
}

// AddWordResponse - model
type AddWordResponse struct {
	BaseResponse
}

// Validate - AddSynonym request and return validation status with proper response
func (awr *AddWordRequest) Validate(r *http.Request) (bool, *AddWordResponse) {

	addWordResponse := new(AddWordResponse)

	// Check if body is empty, because we expect some input
	if r.Body == nil {

		addWordResponse.Code = status.EmptyBody
		addWordResponse.Errors = append(addWordResponse.Errors, Error{Code: status.EmptyBody, Message: status.Text(status.EmptyBody)})
		return false, addWordResponse
	}

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&awr)

	defer r.Body.Close()

	if err != nil {
		addWordResponse.Code = status.IncorrectBodyFormat
		addWordResponse.Errors = append(addWordResponse.Errors, Error{Code: status.IncorrectBodyFormat, Message: status.Text(status.IncorrectBodyFormat)})
		return false, addWordResponse
	}

	log.Printf("awr %v", awr.Word.Word)

	if len(awr.Word.Word) == 0 {
		addWordResponse.Code = status.ErrorMissingWord
		addWordResponse.Errors = append(addWordResponse.Errors, Error{Code: status.ErrorMissingWord, Message: status.Text(status.ErrorMissingWord)})
		return false, addWordResponse
	}

	return true, addWordResponse
}

// Word - a base struct for handling synonyms
type Word struct {
	Word string
}
