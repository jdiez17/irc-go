package irc

import (
	"testing"
	"time"
)

func TestLogIn(t *testing.T) {
	<-time.After(5 * time.Second)

	irc := setUp(t)
	defer irc.Close()

	irc.LogIn(TestIdentity)
}

func TestNewConnection(t *testing.T) {
	<-time.After(5 * time.Second)
	irc := setUp(t)

	err := irc.Close()
	if err != nil {
		t.Error(err)
	}
	if irc.Connection != nil {
		t.Error("Connection isn't nil after Close()")
	}
}
