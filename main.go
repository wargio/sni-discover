package main

import (
	"flag"
	"log"
	"net"
)

type Config struct {
	timeout  int
	target   string
	file     string
	routines int
	best     int
	debug    bool
	h2only   bool
}

func main() {
	var config Config
	var resolver Resolver
	var discover Discover
	var results Results

	flag.BoolVar(&config.debug, "debug", false, "When enabled, shows additional information.")
	flag.BoolVar(&config.h2only, "h2only", false, "When enabled only shows HTTP2 only results.")
	flag.StringVar(&config.target, "target", "", "Target IPv4 address to use for discovering SNIs.")
	flag.StringVar(&config.file, "file", "", "File where to write the SNIs.")
	flag.IntVar(&config.timeout, "timeout", 10, "Connection timeout in seconds.")
	flag.IntVar(&config.routines, "routines", 64, "Maximum number of routines.")
	flag.IntVar(&config.best, "best", 0, "Shows the best N results only when non-zero.")

	flag.Parse()

	if net.ParseIP(config.target) == nil {
		log.Fatal("Invalid -target value")
	} else if config.routines < 1 {
		log.Fatal("Invalid -routines value")
	} else if config.timeout < 1 {
		log.Fatal("Invalid -timeout value")
	} else if config.best < 0 {
		log.Fatal("Invalid -best value")
	}

	discover.Run(&config, &resolver)
	resolver.Run(&config, &results)
	results.Show(&config)
}
