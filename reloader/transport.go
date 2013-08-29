package reloader

import (
	"code.google.com/p/go.crypto/ssh"
	"errors"
	"github.com/mduvall/reloader/reloader/protocols"
	"log"
)

type Transporter interface {
	Publish(string) bool
}

type Transport struct {
	connection *ssh.ClientConn
}

func (t *Transport) Publish(message string) {
	// Some voodoo here to actually send the string over the wire

	log.Println("Got the message: ", message)
}

func CreateTransport(transportParameters map[string]string) (*Transport, error) {
	switch transportParameters["protocol"] {
	case "ssh":
		clientConn, err := reloader.GetSshClientForTransport(transportParameters)
		if err != nil {
			return nil, err
		}
		log.Println("established ssh connection: ", clientConn)

		return &Transport{connection: clientConn}, nil
	}

	return nil, errors.New("not a valid protocol")
}
