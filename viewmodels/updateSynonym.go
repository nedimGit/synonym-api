package viewmodels

import (
	"encoding/json"
	"net/http"

	status "github.com/NedimUka/synonyms/viewmodels/statusCodes"
)

// UpdateSysnonymRequest - model
type UpdateSysnonymRequest struct {
	Word        Word  `json:"Word"`
	UpdatedWord Word  `json:"UpdatedWord"`
	Timestamp   int64 // setup on backend side
}

// UpdateSysnonymResponse - model
type UpdateSysnonymResponse struct {
	BaseResponse
}

// Validate - UpdateSysnonymRequest request and return validation status with proper response
func (usr *UpdateSysnonymRequest) Validate(r *http.Request) (bool, *UpdateSysnonymResponse) {

	updateSysnonymResponse := new(UpdateSysnonymResponse)

	// Check if body is empty, because we expect some input
	if r.Body == nil {

		updateSysnonymResponse.Code = status.EmptyBody
		updateSysnonymResponse.Errors = append(updateSysnonymResponse.Errors, Error{Code: status.EmptyBody, Message: status.Text(status.EmptyBody)})
		return false, updateSysnonymResponse
	}

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&usr)

	defer r.Body.Close()

	if err != nil {
		updateSysnonymResponse.Code = status.IncorrectBodyFormat
		updateSysnonymResponse.Errors = append(updateSysnonymResponse.Errors, Error{Code: status.IncorrectBodyFormat, Message: status.Text(status.IncorrectBodyFormat)})
		return false, updateSysnonymResponse
	}

	if len(usr.Word.Word) == 0 {
		updateSysnonymResponse.Code = status.ErrorMissingWord
		updateSysnonymResponse.Errors = append(updateSysnonymResponse.Errors, Error{Code: status.ErrorMissingWord, Message: status.Text(status.ErrorMissingWord)})
		return false, updateSysnonymResponse
	}

	if len(usr.UpdatedWord.Word) == 0 {
		updateSysnonymResponse.Code = status.ErrorMissingWord
		updateSysnonymResponse.Errors = append(updateSysnonymResponse.Errors, Error{Code: status.ErrorMissingWord, Message: status.Text(status.ErrorMissingWord)})
		return false, updateSysnonymResponse
	}
	return true, updateSysnonymResponse
}
