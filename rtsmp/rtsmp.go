package rtsmp

import (
   "http"
   "net/http"
   "sync"
)

type Server struct {
	mutex   sync.RWMutex
	tls     bool
	http2   bool
}

// Start provides entry point of rtsmp
func Start(configFile string) error {
	config, err := loadConfig(configFile)
	if err != nil {
		return err
	}
}

// Load all servers from config
func startServer(handler bool, addr string) error {
	item, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	if handler {
		item, err := http.ListenAndServe(":8080", nil)
		if err != nil {
			return err
		}
	}
}