package synonym

import (
	"encoding/json"
	"log"
	"net/http"

	synonymService "github.com/NedimUka/synonyms/services"
	vm "github.com/NedimUka/synonyms/viewmodels"
	status "github.com/NedimUka/synonyms/viewmodels/statusCodes"
)

// AddWord - Fucntion that adds new synonym
func AddWord(w http.ResponseWriter, r *http.Request) {

	addWordRequest := new(vm.AddWordRequest)

	valid, addWordResponse := addWordRequest.Validate(r)

	log.Printf("Is valid %v", valid)
	if !valid {
		log.Printf("invalid request received: %v\n", addWordRequest)
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(addWordResponse)
		w.Write(m)
		return
	}

	addedWord, err := synonymService.Instance().AddWords(&addWordRequest.Word)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		m, _ := json.Marshal(addWordResponse)
		w.Write(m)

		return
	}

	addWordResponse.Data = addedWord

	m, err := json.Marshal(addWordResponse)
	if err != nil {
		m, _ := json.Marshal(addWordResponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(m)
		return
	}

	w.Write(m)

	return
}

// GetWords - Fucntion that retreives the list of words
func GetWords(w http.ResponseWriter, r *http.Request) {
	getWordsResponse := new(vm.GetWordsResponse)

	words, err := synonymService.Instance().GetWords()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		m, _ := json.Marshal(getWordsResponse)
		w.Write(m)

		return
	}

	getWordsResponse.Data = words

	m, err := json.Marshal(getWordsResponse)
	if err != nil {
		m, _ := json.Marshal(getWordsResponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(m)
		return
	}

	w.Write(m)

	return
}

// SearchWords - Fucntion that retreives the list of sy
func SearchWords(w http.ResponseWriter, r *http.Request) {

	searchWordsRequest := new(vm.SearchWordsRequest)

	valid, searchWordsresponse := searchWordsRequest.Validate(r)

	if !valid {
		log.Printf("invalid request received: %v\n", searchWordsRequest)
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(searchWordsresponse)
		w.Write(m)
		return
	}

	words, err := synonymService.Instance().SearchWords(searchWordsRequest.Word)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		m, _ := json.Marshal(searchWordsresponse)
		w.Write(m)

		return
	}

	searchWordsresponse.Data = words

	m, err := json.Marshal(searchWordsresponse)
	if err != nil {
		m, _ := json.Marshal(searchWordsresponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(m)
		return
	}

	w.Write(m)

	return
}

// AddSynonym - Add new synonym for existing words
func AddSynonym(w http.ResponseWriter, r *http.Request) {

	addSynonymRequest := new(vm.AddSynonymRequest)

	valid, addSynonymResponse := addSynonymRequest.Validate(r)

	if !valid {
		log.Printf("invalid request received: %v\n", addSynonymRequest)
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(addSynonymResponse)
		w.Write(m)
		return
	}

	addedSynonym, errCode := synonymService.Instance().AddSynonym(addSynonymRequest.Word, addSynonymRequest.Synonym)

	if errCode != 0 {
		switch errCode {
		case status.ErrorWordDoesNotExist:
			addSynonymResponse.Code = int64(errCode)
			addSynonymResponse.Errors = append(
				addSynonymResponse.Errors,
				vm.Error{Code: status.ErrorWordDoesNotExist,
					Message: status.Text(status.ErrorWordDoesNotExist)})
			break
		case status.ErrorSynonymAllreadyEsists:
			addSynonymResponse.Code = int64(errCode)
			addSynonymResponse.Errors = append(addSynonymResponse.Errors,
				vm.Error{Code: status.ErrorSynonymAllreadyEsists,
					Message: status.Text(status.ErrorSynonymAllreadyEsists)})
			break
		}
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(addSynonymResponse)
		w.Write(m)

		return
	}

	addSynonymResponse.Data = addedSynonym

	m, err := json.Marshal(addSynonymResponse)
	if err != nil {
		m, _ := json.Marshal(addSynonymResponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(m)
		return
	}

	w.Write(m)

	return
}

// UpdateSynonym - Update existing synonym
func UpdateSynonym(w http.ResponseWriter, r *http.Request) {

	updateSysnonymRequest := new(vm.UpdateSysnonymRequest)

	valid, updateSynonymResponse := updateSysnonymRequest.Validate(r)

	if !valid {
		log.Printf("invalid request received: %v\n", updateSysnonymRequest)
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(updateSynonymResponse)
		w.Write(m)
		return
	}

	updatedSynonym, errCode := synonymService.Instance().EditWord(updateSysnonymRequest.Word, updateSysnonymRequest.UpdatedWord)

	if errCode != 0 {
		switch errCode {
		case status.ErrorWordDoesNotExist:
			updateSynonymResponse.Code = int64(errCode)
			updateSynonymResponse.Errors = append(
				updateSynonymResponse.Errors,
				vm.Error{Code: status.ErrorWordDoesNotExist,
					Message: status.Text(status.ErrorWordDoesNotExist)})
			break
		case status.ErrorSynonymAllreadyEsists:
			updateSynonymResponse.Code = int64(errCode)
			updateSynonymResponse.Errors = append(updateSynonymResponse.Errors,
				vm.Error{Code: status.ErrorSynonymAllreadyEsists,
					Message: status.Text(status.ErrorSynonymAllreadyEsists)})
			break
		}
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(updateSynonymResponse)
		w.Write(m)

		return
	}

	updateSynonymResponse.Data = updatedSynonym

	m, err := json.Marshal(updateSynonymResponse)
	if err != nil {
		m, _ := json.Marshal(updateSynonymResponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(m)
		return
	}

	w.Write(m)

	return
}

// RemoveSynonym - remove synonym
func RemoveSynonym(w http.ResponseWriter, r *http.Request) {

	removeSynonymRequest := new(vm.RemoveSynonymRequest)

	valid, removeSynonymResponse := removeSynonymRequest.Validate(r)

	if !valid {
		log.Printf("invalid request received: %v\n", removeSynonymRequest)
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(removeSynonymResponse)
		w.Write(m)
		return
	}

	synonymRemoved, errCode := synonymService.Instance().RemoveSynonym(removeSynonymRequest.Word)

	if errCode != 0 {

		removeSynonymResponse.Code = int64(errCode)
		removeSynonymResponse.Errors = append(
			removeSynonymResponse.Errors,
			vm.Error{Code: status.ErrorWordDoesNotExist,
				Message: status.Text(status.ErrorWordDoesNotExist)})

		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(removeSynonymResponse)
		w.Write(m)

		return
	}

	removeSynonymResponse.Data = synonymRemoved

	m, err := json.Marshal(removeSynonymResponse)
	if err != nil {
		m, _ := json.Marshal(removeSynonymResponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(m)
		return
	}

	w.Write(m)

	return
}

// GetSynonyms - list synonyms for specific word
func GetSynonyms(w http.ResponseWriter, r *http.Request) {

	getSynonymsRequest := new(vm.GetSynonymsRequest)

	valid, getSynonymResponse := getSynonymsRequest.Validate(r)

	if !valid {
		log.Printf("invalid request received: %v\n", getSynonymsRequest)
		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(getSynonymsRequest)
		w.Write(m)
		return
	}

	synonyms, errCode := synonymService.Instance().GetSynonyms(getSynonymsRequest.Word)

	if errCode != 0 {

		getSynonymResponse.Code = int64(errCode)
		getSynonymResponse.Errors = append(
			getSynonymResponse.Errors,
			vm.Error{Code: status.ErrorWordDoesNotExist,
				Message: status.Text(status.ErrorWordDoesNotExist)})

		w.WriteHeader(http.StatusBadRequest)
		m, _ := json.Marshal(getSynonymResponse)
		w.Write(m)

		return
	}

	getSynonymResponse.Data = synonyms

	m, err := json.Marshal(getSynonymResponse)
	if err != nil {
		m, _ := json.Marshal(getSynonymResponse)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(m)
		return
	}

	w.Write(m)

	return
}
