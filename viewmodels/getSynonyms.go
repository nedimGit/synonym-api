package viewmodels

import (
	"encoding/json"
	"net/http"

	status "github.com/NedimUka/synonyms/viewmodels/statusCodes"
)

// GetSynonymsRequest - model
type GetSynonymsRequest struct {
	Word      Word  `json:"Word"`
	Timestamp int64 // setup on backend side
}

// GetSynonymsResponse - model
type GetSynonymsResponse struct {
	BaseResponse
}

// Validate - GetSynonyms request and return validation status with proper response
func (awr *GetSynonymsRequest) Validate(r *http.Request) (bool, *GetSynonymsResponse) {

	getSynonymsResponse := new(GetSynonymsResponse)

	// Check if body is empty, because we expect some input
	if r.Body == nil {

		getSynonymsResponse.Code = status.EmptyBody
		getSynonymsResponse.Errors = append(getSynonymsResponse.Errors, Error{Code: status.EmptyBody, Message: status.Text(status.EmptyBody)})
		return false, getSynonymsResponse
	}

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&awr)

	defer r.Body.Close()

	if err != nil {
		getSynonymsResponse.Code = status.IncorrectBodyFormat
		getSynonymsResponse.Errors = append(getSynonymsResponse.Errors, Error{Code: status.IncorrectBodyFormat, Message: status.Text(status.IncorrectBodyFormat)})
		return false, getSynonymsResponse
	}

	if len(awr.Word.Word) == 0 {
		getSynonymsResponse.Code = status.ErrorMissingWord
		getSynonymsResponse.Errors = append(getSynonymsResponse.Errors, Error{Code: status.ErrorMissingWord, Message: status.Text(status.ErrorMissingWord)})
		return false, getSynonymsResponse
	}
	return true, getSynonymsResponse
}
