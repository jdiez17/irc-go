package irc

type EventType int64

const (
	PRIVMSG = iota
	RAW
	MOTD_END
	PING
	MODE
)

type Event struct {
	Type    EventType
	Payload map[string]string
	Params  []string
}

func (e *Event) React(c *Connection, message string) {
	channel := ""
	if e.Payload["channel"] == c.Identity.Nick {
		channel = getNick(e.Payload["sender"])
	} else {
		channel = e.Payload["channel"]
		message = getNick(e.Payload["sender"]) + ": " + message
	}

	c.Privmsg(channel, message)
}
