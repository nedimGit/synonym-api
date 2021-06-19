package viewmodels

import (
	"encoding/json"
	"net/http"

	status "github.com/NedimUka/synonyms/viewmodels/statusCodes"
)

// SearchWordsRequest - model
type SearchWordsRequest struct {
	Word      Word  `json:"Word"`
	Timestamp int64 // setup on backend side
}

// SearchWordsRespnse - model
type SearchWordsRespnse struct {
	BaseResponse
}

// Validate - SearchWords request and return validation status with proper response
func (swr *SearchWordsRequest) Validate(r *http.Request) (bool, *SearchWordsRespnse) {

	searchWordsRespnse := new(SearchWordsRespnse)

	// Check if body is empty, because we expect some input
	if r.Body == nil {

		searchWordsRespnse.Code = status.EmptyBody
		searchWordsRespnse.Errors = append(searchWordsRespnse.Errors, Error{Code: status.EmptyBody, Message: status.Text(status.EmptyBody)})
		return false, searchWordsRespnse
	}

	// Decode request
	err := json.NewDecoder(r.Body).Decode(&swr)

	defer r.Body.Close()

	if err != nil {
		searchWordsRespnse.Code = status.IncorrectBodyFormat
		searchWordsRespnse.Errors = append(searchWordsRespnse.Errors, Error{Code: status.IncorrectBodyFormat, Message: status.Text(status.IncorrectBodyFormat)})
		return false, searchWordsRespnse
	}

	if len(swr.Word.Word) == 0 {
		searchWordsRespnse.Code = status.ErrorMissingWord
		searchWordsRespnse.Errors = append(searchWordsRespnse.Errors, Error{Code: status.ErrorMissingWord, Message: status.Text(status.ErrorMissingWord)})
		return false, searchWordsRespnse
	}
	return true, searchWordsRespnse
}
