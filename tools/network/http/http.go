package http

import (
	"context"
	"log"
	"net/http"

	e "github.com/NedimUka/synonyms/endpoints"
	n "github.com/NedimUka/synonyms/tools/network"
)

// Instance of the service
type Instance struct {
	Configuration n.Config
}

// Setup network
func (i *Instance) Setup(ctx context.Context, configuration n.Config) {
	i.Configuration = configuration

	// Initialize routes
	e.Initialize()
}

// Serve will start serving and listening http reequests
func (i *Instance) Serve() error {
	log.Printf("Serving on port: %v\n", i.Configuration.HTTPPort)

	return http.ListenAndServe(i.Configuration.HTTPPort, nil)
}
