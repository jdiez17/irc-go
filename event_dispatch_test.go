package irc

import (
	"testing"
	"time"
)

func TestEvent376(t *testing.T) {
	<-time.After(5 * time.Second)

	irc := setUp(t)
	defer irc.Close()

	ch := make(chan bool)
	irc.LogIn(TestIdentity)
	irc.AddHandler(MOTD_END, func(c *Connection, e *Event) {
		ch <- true
	})

	select {
	case <-time.After(10 * time.Second):
		t.Error("Timeout while waiting for the 376.")
	case <-ch:
		return
	}
}
