package main

import (
	"fmt"
	"net"

	"github.com/shazow/ssh-chat/chat"
	"github.com/shazow/ssh-chat/sshd"
)

// Identity is a container for everything that identifies a client.
type Identity struct {
	sshd.Connection
	id string
}

// NewIdentity returns a new identity object from an sshd.Connection.
func NewIdentity(conn sshd.Connection) *Identity {
	return &Identity{
		Connection: conn,
		id:         conn.Name(),
	}
}

func (i Identity) Id() string {
	return i.id
}

func (i *Identity) SetId(id string) {
	i.id = id
}

func (i *Identity) SetName(name string) {
	i.SetId(name)
}

func (i Identity) Name() string {
	return i.id
}

func (i Identity) Whois() string {
	ip, _, _ := net.SplitHostPort(i.RemoteAddr().String())
	fingerprint := "(no public key)"
	if i.PublicKey() != nil {
		fingerprint = sshd.Fingerprint(i.PublicKey())
	}
	return fmt.Sprintf("name: %s"+chat.Newline+
		" > ip: %s"+chat.Newline+
		" > fingerprint: %s", i.Name(), ip, fingerprint)
}
