package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/aaronellington/personal-website/assets"
	"github.com/fuzzingbits/forge"
	"github.com/fuzzingbits/philote"
)

// Important paths to be used later
const (
	ContentPath = "./assets/content"
)

// Configuration is the structure of the configuration options
type Configuration struct {
	Host string
	Port string `env:"PORT"`
}

func main() {
	// Get the configuration
	configuration, err := getConfiguration()
	if err != nil {
		panic(err)
	}

	// Setup the server
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", configuration.Host, configuration.Port),
		Handler: getHandler(),
	}

	// Start the server
	log.Printf("Listening on http://%s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func getConfiguration() (*Configuration, error) {
	// Setup defaults
	configuration := &Configuration{
		Host: "0.0.0.0",
		Port: "8000",
	}

	if err := forge.ParseEnvironment(configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}

func getHandler() http.Handler {
	template.New("blah")
	site := &philote.Site{
		Content:  assets.Content,
		Template: assets.Theme,
	}

	// Prime the site
	if err := site.Prime(); err != nil {
		panic(err)
	}

	// Build the primary router
	router := &forge.Router{
		NotFoundHander: site,
	}

	// Configure static file serving
	static := &forge.Static{
		FileSystem:      http.FS(assets.Public),
		NotFoundHandler: router,
	}

	// Configure the logger
	logger := &forge.Logger{
		Handler: static,
	}

	return logger
}
