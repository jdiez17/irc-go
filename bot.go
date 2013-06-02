package irc

type Bot struct {
	Connection *Connection
	Prefix     string
}

func NewBot(c *Connection) Bot {
	return Bot{
		Connection: c,
		Prefix:     ".",
	}
}

func (b *Bot) AddCommand(command string, handler Handler) {
	constraint := Constraint{
		Type:   PRIVMSG,
		Prefix: b.Prefix + command,
	}

	entry := HandlerEntry{
		Handler:         handler,
		EventConstraint: constraint,
	}

	b.Connection.Handlers <- entry
}
