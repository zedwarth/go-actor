package printer

import (
	"fmt"
	"time"
)

type PrinterMessage struct {
	Tick   time.Time
	Number int
}

type PrinterActor struct {
	inbox chan PrinterMessage
}

func NewPrinterActor() *PrinterActor {
	return &PrinterActor{
		inbox: make(chan PrinterMessage, 10),
	}
}

func (p *PrinterActor) Start() {
	go func() {
		for message := range p.inbox {
			fmt.Println("Tick:", message.Tick, "Number:", message.Number)
		}
	}()
}

func (p *PrinterActor) Send(message PrinterMessage) {
	p.inbox <- message
}
