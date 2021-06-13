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

	instance := services.GetSynonymService()

	log.Printf("Adding new word Cleaning")
	instance.AddWords(vm.Word{Word: "Cleaning"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}

	log.Printf("Adding new synonym to Cleaning, Washing")
	instance.AddSynonym(vm.Word{Word: "Cleaning"}, vm.Word{Word: "Washing"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}

	log.Printf("Remove new synonym Cleaning, ")
	instance.RemoveSynonym(vm.Word{Word: "Cleaning"})

	for k, v := range instance.Synonyms {
		log.Printf("Key word %v", k)
		for _, word := range *v {
			log.Printf("Synonym word %v", word.Word)
		}
	}

	// Start server
	server := &http.Instance{}
	server.Setup(context.Background(), httpConfig)

	// Serve and listen
	if err := server.Serve(); err != nil {
		log.Fatalf("Failed to serve and listen: %#v\n", err)
	}

}
