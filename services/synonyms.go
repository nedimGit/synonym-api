package services

import (
	"errors"
	"log"

	vm "github.com/NedimUka/synonyms/viewmodels"
)

// SynonymService - exported pointer to sysnonyms map
type SynonymService struct {
	Synonyms map[string]*[]vm.Word
}

func GetSynonymService() SynonymService {
	return synonymService
}

var synonymService SynonymService

// Init initialises synonym servie
func Init() {
	synonymService = SynonymService{}
	synonyms := make(map[string]*[]vm.Word)
	synonymService.Synonyms = synonyms
}

// AddWords - Add new word without synonyms
func (service *SynonymService) AddWords(word vm.Word) (*[]vm.Word, error) {

	// Check if the word already exists
	if val, ok := service.Synonyms[word.Word]; ok {
		return val, nil

	}

	// If the word did not exist before register it and return the empty initialised slice of words
	words := []vm.Word{word}
	service.Synonyms[word.Word] = &words
	return &words, nil

}

// AddSynonym - Add new synonym to word
func (service *SynonymService) AddSynonym(word vm.Word, synonym vm.Word) (*[]vm.Word, error) {

	// Check if synonym already exists
	if _, ok := service.Synonyms[synonym.Word]; ok {
		return nil, errors.New("The word alreeay exists as a synonym")
	}

	// Check if the word already exists
	if val, ok := service.Synonyms[word.Word]; ok {

		log.Printf("val %v", val)
		log.Printf("val value %v", *val)
		*val = append(*val, synonym)
		service.Synonyms[synonym.Word] = val
		return val, nil

	}

	return nil, errors.New("The word and its synonyms does not exist")

}

// EditWord - Edit existing word
func (service *SynonymService) EditWord(word vm.Word, replacement vm.Word) (*[]vm.Word, error) {

	//Check if new word aready exists
	if _, ok := service.Synonyms[replacement.Word]; ok {
		return nil, errors.New("Word already exists")
	}

	// Check if word already exists
	if val, ok := service.Synonyms[word.Word]; ok {

		for _, word := range *val {
			existingWord := &word
			existingWord.Word = replacement.Word
		}
		service.Synonyms[replacement.Word] = val
		delete(service.Synonyms, word.Word)

		return val, nil

	}

	return nil, errors.New("The word and its synonyms does not exist")
}

// RemoveSynonym - Remove existing synonym
func (service *SynonymService) RemoveSynonym(synonym vm.Word) (bool, error) {

	// Check if synonym already exists
	if val, ok := service.Synonyms[synonym.Word]; ok {

		for i, word := range *val {
			if word.Word == synonym.Word {
				*val = append((*val)[:i], (*val)[i+1:]...)
			}
		}
		delete(service.Synonyms, synonym.Word)
		return true, nil

	}

	return false, errors.New("The word and its synonyms does not exist")

}
