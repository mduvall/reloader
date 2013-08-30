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
	Connection *ssh.ClientConn
}

/**
 * Establishes a session with the ClientConn and starts the netcat session
 * for RPC
 */
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

		return &Transport{Connection: clientConn}, nil
	}

	return nil, errors.New("not a valid protocol")
}
