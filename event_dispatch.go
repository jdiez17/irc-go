package irc

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func reader(conn net.Conn, ch chan<- string, close chan bool) {
	reader := bufio.NewReader(conn)

	for {
		ev, err := reader.ReadString('\n')
		if err != nil {
			close <- true
			return
		}

		ch <- ev
	}
}

func generateEvent(data string) Event {
	var t EventType = RAW
	var sender string
	var command string

	idx := 0
	if data[idx] == ':' {
		idx++
		for data[idx] != ' ' {
			sender += string(data[idx])
			idx++
		}

		idx++
		for data[idx] != ' ' {
			command += string(data[idx])
			idx++
		}

		switch command {
		case "376":
			t = MOTD_END

		case "PRIVMSG":
			t = PRIVMSG

			var channel string
			var message string

			idx++
			for data[idx] != ' ' {
				channel += string(data[idx])
				idx++
			}

			idx += 2
			for idx < len(data) {
				message += string(data[idx])
				idx++
			}

			payload := map[string]string{
				"sender":  sender,
				"channel": channel,
				"message": message,
				"raw":     data,
			}

			return Event{Type: t, Payload: payload}
		}
	}
	parts := strings.Split(data, " ")
	if parts[0] == "PING" {
		return Event{Type: PING, Payload: map[string]string{
			"raw":      data,
			"response": parts[1],
		}}
	}

	return Event{Type: t, Payload: map[string]string{"raw": data}}
}

func eventDispatcher(conn *Connection, ch chan<- Event, close chan bool) {
	dataChan := make(chan string)
	go reader(conn.Connection, dataChan, close)

	for {

		select {
		case <-close:
			break
		case data := <-dataChan:
			log.Println("<<< " + data)
			ch <- generateEvent(data)
		}
	}
	conn.Close()
}
