package irc

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Identity struct {
	Nick string
}

type Connection struct {
	Connection net.Conn
	Events     chan Event
	Shutdown   chan bool
	Handlers   chan HandlerEntry

	Identity Identity
}

type connectionRequest struct {
	Server     string
	Port       int
	Connection net.Conn
}

func NewConnection(server string, port int) (conn *Connection, err error) {
	return createConnection(connectionRequest{Server: server, Port: port})
}

func NewConnectionCustomConn(c net.Conn) (conn *Connection, err error) {
	return createConnection(connectionRequest{Connection: c})
}

func createConnection(rq connectionRequest) (conn *Connection, err error) {
	var sock net.Conn
	if rq.Server != "" && rq.Port != 0 {
		sock, err = net.Dial("tcp", fmt.Sprintf("%s:%d", rq.Server, rq.Port))
	} else {
		sock = rq.Connection
	}

	ch := make(chan Event)
	shutdown := make(chan bool)
	handlers := make(chan HandlerEntry)

	conn = &Connection{Connection: sock, Events: ch, Shutdown: shutdown, Handlers: handlers}
	go eventDispatcher(conn, ch, shutdown)
	go routeEvents(conn, ch, handlers)

	conn.AddHandler(PING, func(c *Connection, e *Event) {
		c.Write("PONG " + e.Payload["response"])
	})

	return
}

func (c *Connection) Close() (err error) {
	c.Quit("irc.Close()")

	c.Shutdown <- true

	err = c.Connection.Close()
	c.Connection = nil
	return
}

func (c *Connection) Write(s string) (n int, err error) {
	s = strings.Replace(s, "\n", "", -1)
	log.Println(">>> " + s)
	return c.Connection.Write([]byte(s + "\n"))
}

func (c *Connection) AddHandler(typ EventType, handler Handler) {
	constraint := Constraint{Type: typ}
	entry := HandlerEntry{handler, constraint}

	c.Handlers <- entry
}
