package irc

import (
	"fmt"
)

func (c *Connection) Password(password string) {
	c.Write(fmt.Sprintf(C_PASS, password))
}

func (c *Connection) LogInPassword(i Identity, password string) {
	c.Identity = i

	c.Write(fmt.Sprintf(C_PASS, password))
	c.Write(fmt.Sprintf(C_USER, i.Nick, i.Nick, i.Nick, i.Nick))
	c.Write(fmt.Sprintf(C_NICK, i.Nick))
}

func (c *Connection) LogIn(i Identity) {
	c.Identity = i

	c.Write(fmt.Sprintf(C_USER, i.Nick, i.Nick, i.Nick, i.Nick))
	c.Write(fmt.Sprintf(C_NICK, i.Nick))
}

func (c *Connection) Quit(message string) {
	c.Write(fmt.Sprintf(C_QUIT, message))
}

func (c *Connection) Join(channel string) {
	c.Write(fmt.Sprintf(C_JOIN, channel))
}

func (c *Connection) Privmsg(target, message string) {
	c.Write(fmt.Sprintf(C_PRIVMSG, target, message))
}

func (c *Connection) Part(channel string) {
	c.Write(fmt.Sprintf(C_PART, channel))
}
