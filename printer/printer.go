package printer

import (
	"fmt"
	"time"
)

type Message struct {
	Tick   time.Time
	Number int
}

type Actor struct {
	inbox chan Message
}

func NewActor(buffer int) *Actor {
	return &Actor{
		inbox: make(chan Message, buffer),
	}
}

func (a *Actor) Start() {
	go func() {
		for message := range a.inbox {
			handler(message)
		}
	}()
}

func (a *Actor) Send(message Message) {
	a.inbox <- message
}

func handler(message Message) {
	fmt.Println("Tick:", message.Tick, "Number:", message.Number)
}
