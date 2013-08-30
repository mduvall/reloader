package main

import (
	"github.com/mduvall/reloader/reloader"
	"log"
)

func main() {

	done := make(chan bool)

	// TODO: pull this into a config file

	client, err := reloader.CreateTransport(map[string]string{})
	server, err := reloader.CreateServer()
	if err != nil {
		log.Fatal(err)
	}

	server.ListenAtPath("/")

	<-done
}
