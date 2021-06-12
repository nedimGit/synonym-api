package synonym

import (
	"net/http"

	vm "github.com/NedimUka/synonyms/viewmodels"
)

// Add - Fucntion that adds new synonym
func Add(w http.ResponseWriter, r *http.Request) {

	addSynonymRequest := new(vm.AddSynonymRequest)

	valid, addSynonymResponse := addSynonymRequest.Validate(r)

	return tru

}
