package irc

import "testing"

var TestIdentity Identity = Identity{Nick: "goLangTestBot"}

func setUp(t *testing.T) *Connection {
	conn, err := NewConnection("irc.freenode.net", 6667)
	if err != nil {
		t.Error(err)
	}

	return conn
}
