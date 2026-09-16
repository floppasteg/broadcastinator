package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Writer()
	masconfPath := flag.String("c", "", "location of master config")
	noweb := flag.Bool("noweb", false, "don't load web interface")

	help := flag.Bool("help", false, "show long help (recommended)")
	flag.Parse()

	if *help {
		fmt.Fprintln(os.Stderr, `broadcastinator: youtube-oriented RTMP broadcast thingy

Options:
	-c <configfile>   path to the configuration file
	--noweb           don't load web interface
	--render          render out the entire channel into a single video file
ni`)
	}
	if !*noweb {
		fmt.Fprintln(os.Stderr, "Loading web interface")
		go loadInterface()
	}
	fmt.Fprintln(os.Stderr, "Loading API")
	go loadAPI()
	masconf, err := os.Open(*masconfPath)
	chk(err)
	config := MasterConfig{}
	config = parseGeneric(masconf, config)

	// now we wait...
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
