package config

import "os"

// ServerConfigInterface server config interface
type ServerConfigInterface interface {
	Port() string
	Mode() string
}

// Server server config struct
type Server struct {
	port string
	mode string
}

// NewServerConfig create server config struct instance
func NewServerConfig() *Server {
	port := "4000"
	mode := "debug"

	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("MODE"); env != "" {
		mode = env
	}

	server := &Server{
		mode: mode,
		port: port,
	}

	if server.mode != "release" && server.mode != "debug" {
		panic("Unavailable gin mode")
	}

	return server
}

// Port get server port number
func (server *Server) Port() string {
	return server.port
}

// Mode get server mode
func (server *Server) Mode() string {
	return server.mode
}
