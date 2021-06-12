package models

// Config stores application configuration
type Config struct {
	Service Service
}

// Service contains configuration for http service
type Service struct {
	Port        string
	HTTPTimeout int
}
