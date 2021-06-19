package viewmodels

import (
	"encoding/json"
	"net/http"

	status "github.com/NedimUka/synonyms/viewmodels/statusCodes"
)

// AddSynonymRequest - model
type AddSynonymRequest struct {
	Word      Word  `json:"Word"`
	Synonym   Word  `json:"Synonym"`
	Timestamp int64 // setup on backend side
}

// AddSynonymResponse - model
type AddSynonymResponse struct {
	BaseResponse
}

// Validate - AddSynonym request and return validation status with proper response
func (asr *AddSynonymRequest) Validate(r *http.Request) (bool, *AddSynonymResponse) {

	addSynonymResponse := new(AddSynonymResponse)

	// Check if body is empty, because we expect some input
	if r.Body == nil {

		addSynonymResponse.Code = status.EmptyBody
		addSynonymResponse.Errors = append(addSynonymResponse.Errors, Error{Code: status.EmptyBody, Message: status.Text(status.EmptyBody)})
		return false, addSynonymResponse
	}

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&asr)

	defer r.Body.Close()

	if err != nil {
		addSynonymResponse.Code = status.IncorrectBodyFormat
		addSynonymResponse.Errors = append(addSynonymResponse.Errors, Error{Code: status.IncorrectBodyFormat, Message: status.Text(status.IncorrectBodyFormat)})
		return false, addSynonymResponse
	}

	if len(asr.Word.Word) == 0 {
		addSynonymResponse.Code = status.ErrorMissingWord
		addSynonymResponse.Errors = append(addSynonymResponse.Errors, Error{Code: status.ErrorMissingWord, Message: status.Text(status.ErrorMissingWord)})
		return false, addSynonymResponse
	}

	if len(asr.Synonym.Word) == 0 {
		addSynonymResponse.Code = status.ErrorMissingWord
		addSynonymResponse.Errors = append(addSynonymResponse.Errors, Error{Code: status.ErrorMissingWord, Message: status.Text(status.ErrorMissingWord)})
		return false, addSynonymResponse
	}

	return true, addSynonymResponse
}
