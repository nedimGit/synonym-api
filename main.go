package main

import (
	"context"
	"log"
	"time"

	"github.com/NedimUka/synonyms/config"
	"github.com/NedimUka/synonyms/tools/network"
	"github.com/NedimUka/synonyms/tools/network/http"
)

func main() {

	config.Load()

	// Configure the server
	httpConfig := network.Config{
		HTTPPort:         ":" + config.AppConfig.Service.Port,
		HTTPReadTimeout:  time.Duration(config.AppConfig.Service.HTTPTimeout) * time.Second,
		HTTPWriteTimeout: time.Duration(config.AppConfig.Service.HTTPTimeout) * time.Second,
	}

	// Start server
	server := &http.Instance{}
	server.Setup(context.Background(), httpConfig)

	// Serve and listen
	if err := server.Serve(); err != nil {
		log.Fatalf("Failed to serve and listen: %#v\n", err)
	}
}
