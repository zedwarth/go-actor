package generator

import (
	"math/rand"
	"time"

	"github.com/zedwarth/go-actor/printer"
)

type Message struct {
	Tick    time.Time
	Printer *printer.Actor
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
	message.Printer.Send(printer.Message{
		Tick:   message.Tick,
		Number: numberGenerator(),
	})
}

func numberGenerator() int {
	mean := 50.0
	stddev := 100.0
	n := int(rand.NormFloat64()*stddev + mean)
	return n
}
