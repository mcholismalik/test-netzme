package main

import (
	"github.com/test-netzme/internal/config"
	"github.com/test-netzme/internal/driver/http"
	"github.com/test-netzme/internal/factory"
)

// @title test-netzme
// @version 0.0.1
// @description This is a doc for test-netzme.

// @host localhost:3000
// @BasePath /
func main() {
	cfg, err := config.Load("")
	if err != nil {
		panic(err)
	}

	// factory
	f, err := factory.Init(cfg)
	if err != nil {
		panic(err)
	}

	// http
	http.Init(cfg, f)
}
