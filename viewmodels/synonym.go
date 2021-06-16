package viewmodels

import "net/http"

// AddSynonymRequest - model
type AddSynonymRequest struct {
	Synonyms  []Word `json:"Synonyms"`
	Timestamp int64  // setup on backend side
}

// AddSynonymResponse - model
type AddSynonymResponse struct {
	BaseResponse
}

// Validate - AddSynonym request and return validation status with proper response
func (sM AddSynonymRequest) Validate(r *http.Request) (bool, *AddSynonymResponse) {

	addSynonymResponse := new(AddSynonymResponse)

	// // Check if body is empty, because we expect some input
	// if r.Body == nil {
	// 	createMeasureResponse.Code = status.EmptyBody
	// 	createMeasureResponse.Errors = append(createMeasureResponse.Errors, Error{Code: status.EmptyBody, Message: status.Text(status.EmptyBody)})
	// 	return false, createMeasureResponse
	// }

	// // Decode request
	// err := json.NewDecoder(r.Body).Decode(&cM)

	// defer r.Body.Close()

	// if err != nil {
	// 	createMeasureResponse.Code = status.IncorrectBodyFormat
	// 	createMeasureResponse.Errors = append(createMeasureResponse.Errors, Error{Code: status.IncorrectBodyFormat, Message: status.Text(status.IncorrectBodyFormat)})
	// 	return false, createMeasureResponse
	// }
	return true, addSynonymResponse
}
