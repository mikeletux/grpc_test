package server

// Server define a server behavior
type Server interface {
	// Serve servers a service's server implementation
	Serve() error
}

// Config contains the configuration to set up a server
type Config struct {
	Protocol string //TCP UDP
	Host     string
	Port     string
}
