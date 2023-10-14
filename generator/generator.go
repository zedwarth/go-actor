package generator

import (
	"math"
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
		Number: numberGenerator(50, 15),
	})
}

func numberGenerator(mean float64, standardDeviation float64) int {
	value := rand.NormFloat64()*standardDeviation + mean
	value = math.Max(0, math.Min(100, value))

	return int(value)
}
