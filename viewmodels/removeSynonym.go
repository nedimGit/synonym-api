package viewmodels

import (
	"encoding/json"
	"net/http"

	status "github.com/NedimUka/synonyms/viewmodels/statusCodes"
)

// RemoveSynonymRequest - model
type RemoveSynonymRequest struct {
	Word      Word  `json:"Word"`
	Timestamp int64 // setup on backend side
}

// RemoveSynonymResponse - model
type RemoveSynonymResponse struct {
	BaseResponse
}

// Validate - RemoveSynonym request and return validation status with proper response
func (rsr *RemoveSynonymRequest) Validate(r *http.Request) (bool, *RemoveSynonymResponse) {

	removeSynonymResponse := new(RemoveSynonymResponse)

	// Check if body is empty, because we expect some input
	if r.Body == nil {

		removeSynonymResponse.Code = status.EmptyBody
		removeSynonymResponse.Errors = append(removeSynonymResponse.Errors, Error{Code: status.EmptyBody, Message: status.Text(status.EmptyBody)})
		return false, removeSynonymResponse
	}

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&rsr)

	defer r.Body.Close()

	if err != nil {
		removeSynonymResponse.Code = status.IncorrectBodyFormat
		removeSynonymResponse.Errors = append(removeSynonymResponse.Errors, Error{Code: status.IncorrectBodyFormat, Message: status.Text(status.IncorrectBodyFormat)})
		return false, removeSynonymResponse
	}

	if len(rsr.Word.Word) == 0 {
		removeSynonymResponse.Code = status.ErrorMissingWord
		removeSynonymResponse.Errors = append(removeSynonymResponse.Errors, Error{Code: status.ErrorMissingWord, Message: status.Text(status.ErrorMissingWord)})
		return false, removeSynonymResponse
	}
	return true, removeSynonymResponse
}
