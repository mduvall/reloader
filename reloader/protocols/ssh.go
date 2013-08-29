package reloader

import (
	"code.google.com/p/go.crypto/ssh"
	"errors"
	"log"
)

type ClientPassword string

func (p ClientPassword) Password(user string) (string, error) {
	return string(p), nil
}

// Just lob over the password when questioned
type PasswordInteractive string

func (p PasswordInteractive) Challenge(user, instruction string, questions []string, echos []bool) ([]string, error) {
	passwords := make([]string, len(questions))
	for i, _ := range passwords {
		passwords[i] = string(p)
	}

	return passwords, nil
}

func GetSshClientForTransport(sshParameters map[string]string) (*ssh.ClientConn, error) {
	conf := &ssh.ClientConfig{
		User: sshParameters["user"],
		Auth: []ssh.ClientAuth{
			ssh.ClientAuthKeyboardInteractive(
				PasswordInteractive(sshParameters["password"])),
			// We can try password auth here, usually is just keyboard-interactive/publickey though
			ssh.ClientAuthPassword(ClientPassword(sshParameters["password"])),
		},
	}

	client, err := ssh.Dial("tcp", sshParameters["host"]+":"+sshParameters["port"], conf)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to create ssh connection")
	}

	return client, nil
}
