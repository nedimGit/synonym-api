package synonym

import (
	"encoding/json"
	"log"
	"net/http"

	vm "github.com/NedimUka/synonyms/viewmodels"
)

// Add - Fucntion that adds new synonym
func Add(w http.ResponseWriter, r *http.Request) {

	// addSynonymRequest := new(vm.AddSynonymRequest)

	// valid, addSynonymResponse := addSynonymRequest.Validate(r)

	// return tru

}

// AddWord - Fucntion that adds new synonym
func AddWord(w http.ResponseWriter, r *http.Request) {

	addWordRequest := new(vm.AddWordRequest)

	valid, addWordResponse := addWordRequest.Validate(r)

	if !valid {
		log.Printf("")
	}

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
