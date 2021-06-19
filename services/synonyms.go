package services

import (
	"log"
	"regexp"

	vm "github.com/NedimUka/synonyms/viewmodels"
	status "github.com/NedimUka/synonyms/viewmodels/statusCodes"
)

// SynonymService - exported pointer to sysnonyms map
type SynonymService struct {
	Synonyms map[string]*[]vm.Word
}

// Instance - get instance of synonym service
func Instance() *SynonymService {
	return &synonymService
}

var synonymService SynonymService

// Init initialises synonym servie
func Init() {
	synonymService = SynonymService{}
	synonyms := make(map[string]*[]vm.Word)
	synonymService.Synonyms = synonyms
}

// AddWords - Add new word without synonyms
func (service *SynonymService) AddWords(word *vm.Word) (*[]vm.Word, error) {

	// Check if the word already exists
	if val, ok := service.Synonyms[word.Word]; ok {
		return val, nil

	}

	// If the word did not exist before register it and return the empty initialised slice of words
	words := []vm.Word{*word}
	service.Synonyms[word.Word] = &words
	return &words, nil

}

// AddSynonym - Add new synonym to word
func (service *SynonymService) AddSynonym(word vm.Word, synonym vm.Word) (*[]vm.Word, int) {

	// Check if synonym already exists
	if _, ok := service.Synonyms[synonym.Word]; ok {
		return nil, status.ErrorSynonymAllreadyEsists
	}

	// Check if the word already exists
	if val, ok := service.Synonyms[word.Word]; ok {

		*val = append(*val, synonym)
		service.Synonyms[synonym.Word] = val
		return val, 0

	}

	return nil, status.ErrorWordDoesNotExist

}

// EditWord - Edit existing word
func (service *SynonymService) EditWord(word vm.Word, replacement vm.Word) (*[]vm.Word, int) {

	//Check if new word aready exists
	if _, ok := service.Synonyms[replacement.Word]; ok {
		return nil, status.ErrorSynonymAllreadyEsists
	}

	// Check if word already exists
	if val, ok := service.Synonyms[word.Word]; ok {

		for i, existingWord := range *val {

			if existingWord.Word == word.Word {
				(*val)[i].Word = replacement.Word
			}
		}
		service.Synonyms[replacement.Word] = val
		delete(service.Synonyms, word.Word)

		return val, 0

	}

	return nil, status.ErrorWordDoesNotExist
}

// RemoveSynonym - Remove existing synonym
func (service *SynonymService) RemoveSynonym(synonym vm.Word) (bool, int) {

	// Check if synonym already exists
	if val, ok := service.Synonyms[synonym.Word]; ok {

		for i, word := range *val {
			if word.Word == synonym.Word {
				*val = append((*val)[:i], (*val)[i+1:]...)
			}
		}
		delete(service.Synonyms, synonym.Word)
		return true, 0

	}

	return false, status.ErrorWordDoesNotExist

}

// SearchWords - Search all words
func (service *SynonymService) SearchWords(searchTerm vm.Word) ([]vm.Word, error) {

	words := new([]vm.Word)
	for key := range service.Synonyms {

		match, err := regexp.Match(".*"+searchTerm.Word+".*", []byte(key))

		if err != nil {
			log.Printf("Unable to match word %v error %v ", searchTerm.Word, err)
			continue
		}

		if match {
			matchedWord := vm.Word{Word: key}
			*words = append(*words, matchedWord)
		}

	}

	return *words, nil

}

// GetSynonyms - Get synonyms for specific word
func (service *SynonymService) GetSynonyms(word vm.Word) ([]vm.Word, int) {

	// Check if word already exists
	if val, ok := service.Synonyms[word.Word]; ok {

		return *val, 0

	}

	return nil, status.ErrorWordDoesNotExist

}

// GetWords - Get all words
func (service *SynonymService) GetWords() ([]vm.Word, error) {

	words := make([]vm.Word, 0)
	// Check if word already exists
	for k := range service.Synonyms {
		words = append(words, vm.Word{Word: k})
	}
	return words, nil

}
