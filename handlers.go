package irc

import (
	"strings"
	"text/scanner"
)

type Handler func(c *Connection, e *Event)

type Constraint struct {
	Type   EventType
	Prefix string
}

type HandlerEntry struct {
	Handler         Handler
	EventConstraint Constraint
}

func getParameters(command string) []string {
	var result []string
	var s scanner.Scanner
	s.Init(strings.NewReader(command))

	tok := s.Scan()
	for tok != scanner.EOF {
		text := s.TokenText()
		text = strings.Replace(text, "\"", "", -1)

		result = append(result, text)
		tok = s.Scan()
	}

	return result
}

func checkConstraintSatisfied(c Constraint, e *Event) bool {
	if e.Type == c.Type {
		if c.Prefix != "" {
			if strings.Index(e.Payload["message"], c.Prefix) == 0 {
				params := strings.Replace(e.Payload["message"], c.Prefix, "", 1)
				e.Params = getParameters(params)

				params_len := len(e.Params)
				if params_len >= 2 {
					if e.Params[params_len-2] == "@" {
						e.Payload["sender"] = e.Params[params_len-1]
						e.Params = e.Params[:params_len-2]
					}
				}
				return true
			}
			return false
		}
		return true
	}
	return false
}

func routeEvents(c *Connection, events chan Event, rqs chan HandlerEntry) {
	mapping := make([]HandlerEntry, 0)

	for {
		select {
		case entry := <-rqs:
			mapping = append(mapping, entry)
		case event := <-events:
			for _, i := range mapping {
				if checkConstraintSatisfied(i.EventConstraint, &event) {
					go func(h Handler, e *Event) {
						defer func() {
							r := recover()
							for r != nil {
								err, _ := r.(error)
								e.React(c, "Recovered from panic: "+err.Error())

								r = recover()
							}
						}()

						h(c, &event)
					}(i.Handler, &event)
				}
			}
		}
	}
}
