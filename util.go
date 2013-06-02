package irc

import "strings"

func getNick(uid string) string {
	return strings.Split(uid, "!")[0]
}
