package main

import (
	"context"
	"log"
	"time"

	"github.com/NedimUka/synonyms/config"
	"github.com/NedimUka/synonyms/services"
	"github.com/NedimUka/synonyms/tools/network"
	"github.com/NedimUka/synonyms/tools/network/http"
	vm "github.com/NedimUka/synonyms/viewmodels"
)

func main() {

	config.Load()

	// Configure the server
	httpConfig := network.Config{
		HTTPPort:         ":" + config.AppConfig.Service.Port,
		HTTPReadTimeout:  time.Duration(config.AppConfig.Service.HTTPTimeout) * time.Second,
		HTTPWriteTimeout: time.Duration(config.AppConfig.Service.HTTPTimeout) * time.Second,
	}

	services.Init()

	instance := services.Instance()

	log.Printf("**************************************************************")
	log.Printf("Adding new word Cleaning")
	instance.AddWords(&vm.Word{Word: "Cleaning"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}
	log.Printf("**************************************************************")
	log.Printf("Adding Duplicate  word Cleaning")
	instance.AddWords(&vm.Word{Word: "Cleaning"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}
	log.Printf("**************************************************************")
	log.Printf("Adding new synonym to Cleaning, Washing")
	instance.AddSynonym(vm.Word{Word: "Cleaning"}, vm.Word{Word: "Washing"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}
	log.Printf("**************************************************************")

	log.Printf("Remove new synonym Cleaning, ")
	instance.RemoveSynonym(vm.Word{Word: "Cleaning"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}

	log.Printf("**************************************************************")
	log.Printf("Adding new synonym to Washing, Cleaning")
	instance.AddSynonym(vm.Word{Word: "Washing"}, vm.Word{Word: "Cleaning"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}

	log.Printf("**************************************************************")
	log.Printf("Editng to  Washing to Wash")
	instance.EditWord(vm.Word{Word: "Washing"}, vm.Word{Word: "Wash"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}

	log.Printf("**************************************************************")
	log.Printf("Search word containing letters (was)")
	words, _ := instance.SearchWords(vm.Word{Word: "Was"})

	for _, word := range words {
		log.Printf("Word found %v", word)

	}

	log.Printf("**************************************************************")
	log.Printf("Search all words")
	words, _ = instance.SearchWords(vm.Word{Word: ""})

	for _, word := range words {
		log.Printf("Word found %v", word)

	}

	log.Printf("**************************************************************")
	log.Printf("Adding new word Scrubing")
	instance.AddWords(&vm.Word{Word: "Scrubing"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}

	log.Printf("**************************************************************")
	log.Printf("Search all words that contains (ing)")
	words, _ = instance.SearchWords(vm.Word{Word: "ing"})

	for _, word := range words {
		log.Printf("Word found %v", word)

	}

	log.Printf("**************************************************************")
	log.Printf("Get all words for word ing")
	words, _ = instance.GetSynonyms(vm.Word{Word: "ing"})

	log.Printf("Synonnym for ing found %v", &words)

	log.Printf("**************************************************************")
	log.Printf("Get all words that of wodrd (Cleaning)")
	words, _ = instance.GetSynonyms(vm.Word{Word: "Cleaning"})

	for _, word := range words {
		log.Printf("Word found %v", word)

	}
	// Start server
	server := &http.Instance{}
	server.Setup(context.Background(), httpConfig)

	// Serve and listen
	if err := server.Serve(); err != nil {
		log.Fatalf("Failed to serve and listen: %#v\n", err)
	}

}
