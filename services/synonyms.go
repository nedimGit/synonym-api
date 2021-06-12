package services

import (
	vm "github.com/NedimUka/synonyms/viewmodels"
)

// Instance - exported pointer to sysnonyms map
var Instance *map[string][]*vm.Word

//  sysnonyms map that contains word and pointer to slice of that word's synonyms
var synonyms map[string][]*vm.Word

// Init initialises synonym servie
func Init() {

	synonyms = make(map[string][]*vm.Word)
	Instance = &synonyms
}
