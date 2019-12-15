package main

import (
	"flag"
	"fmt"
	"os"

	"urlshortener/miscellaneous"
	"urlshortener/server"
)

// Global vars for versioning
var (
	Build     = "unset" // nolint
	BuildDate = "unset" // nolint
	GoVersion = "unset" // nolint
	Version   = "unset" // nolint
)

func main() {
	var (
		bind         = flag.String("bind", "0.0.0.0:8000", "Address to bind to, host:port")
		readTimeout  = flag.Int("readt", 30, "Read timeout, seconds")
		writeTimeout = flag.Int("writet", 30, "Write timeout, seconds")
		urlDBPath    = flag.String("urldb", "", "Path to URLDB CSV file")
		templateDir  = flag.String("tpldir", "templates", "Path to templates")
		version      = flag.Bool("version", false, "Print version and exit")
	)

	flag.Parse()

	if *version {
		fmt.Printf("%s version %s\n", server.ProductName, Version)
		fmt.Printf("Build: %s\n", Build)
		fmt.Printf("BuildDate: %s\n", BuildDate)
		fmt.Printf("Go: %s\n\n", GoVersion)
		os.Exit(0)
	}

	serverConfig := &miscellaneous.ServerConfig{
		BuildInfo: &miscellaneous.BuildInfo{
			Build:     Build,
			BuildDate: BuildDate,
			GoVersion: GoVersion,
			Version:   Version,
		},
		Bind:              *bind,
		ReadTimeout:       *readTimeout,
		WriteTimeout:      *writeTimeout,
		URLDBFile:         *urlDBPath,
		TemplateDirectory: *templateDir,
	}

	server.Run(serverConfig)
}
