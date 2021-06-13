package services

import (
	"errors"

	vm "github.com/NedimUka/synonyms/viewmodels"
)

// SynonymService - exported pointer to sysnonyms map
type SynonymService struct {
	Synonyms map[string]*[]*vm.Word
}

var synonymService SynonymService

// Init initialises synonym servie
func Init() {
	synonymService = SynonymService{}
	synonyms := make(map[string]*[]*vm.Word)
	synonymService.Synonyms = synonyms
}

// AddWords - Add new word to synonmys
func (service *SynonymService) AddWords(word vm.Word) (*[]*vm.Word, error) {

	// Check if the word already exists
	if val, ok := service.Synonyms[word.Word]; ok {
		return val, nil

	}

	// If the word did not exist before register it and return the empty initialised slice

	words := []*vm.Word{&word}
	service.Synonyms[word.Word] = &words
	return &words, nil

}

// AddSynonym - Add new word to synonmys
func (service *SynonymService) AddSynonym(word vm.Word, synonym vm.Word) (*[]*vm.Word, error) {

	if _, ok := service.Synonyms[synonym.Word]; ok {

		return nil, errors.New("The word allreedy exists as a synonym")

	}

	// Check if the word already exists
	if val, ok := service.Synonyms[word.Word]; ok {

		*val = append(*val, &synonym)
		*service.Synonyms[synonym.Word] = *val
		return val, nil

	}

	return nil, errors.New("The word and its synonyms does not exist")

}
