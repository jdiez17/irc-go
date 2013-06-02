package irc

import (
	"testing"
	"time"
)

func TestJoin(t *testing.T) {
	<-time.After(5 * time.Second)

	irc := setUp(t)
	defer irc.Close()

	irc.LogIn(TestIdentity)
	irc.AddHandler(MOTD_END, func(c *Connection, e *Event) {
		irc.Join("#botwar")
	})

	<-time.After(30 * time.Second)
}
