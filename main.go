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
	noweb := flag.Bool("noweb", false, "don't load web interface")
	flag.Parse()
	if !*noweb {
		fmt.Fprintln(os.Stderr, "Loading web interface")
		go loadInterface()
		fmt.Fprintln(os.Stderr, "Loading API")
		go loadAPI()
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
