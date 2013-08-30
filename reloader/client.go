package reloader

import (
	"fmt"
	"io"
	"log"
)

type Client struct {
	transport *Transport
	inPipe    io.Writer
	outPipe   io.Reader
}

func CreateClient(transportParams map[string]string) (*Client, error) {
	transport, err := CreateTransport(map[string]string{})

	if err != nil {
		log.Println("failed to create transport for ", transportParams["protocol"])
		return nil, err
	}

	in, out, err := netcatPath(transport, "/tmp/rl.sock")
	if err != nil {
		return nil, err
	}

	return &Client{transport: transport, inPipe: in, outPipe: out}, nil
}

func (c *Client) Message(message string) {
	c.transport.Publish(message)
}

func netcatPath(transport *Transport, path string) (io.Writer, io.Reader, error) {
	session, err := transport.Connection.NewSession()

	if err != nil {
		return nil, nil, err
	}

	cmd := fmt.Sprintf("nc -U %s", path)

	in, err := session.StdinPipe()
	if err != nil {
		return nil, nil, err
	}

	out, err := session.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	err = session.Start(cmd)
	if err != nil {
		return nil, nil, err
	}

	return in, out, nil
}
