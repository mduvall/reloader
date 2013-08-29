package reloader

import (
	"log"
)

type Transporter interface {
	Publish(string) bool
}

type Transport struct {
	protocol string
}

func (t *Transport) Publish(message string) {
	// Some voodoo here to actually send the string over the wire

	log.Println("Got the message: ", message)
}
