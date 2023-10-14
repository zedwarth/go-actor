package main

import (
	"time"

	"github.com/zedwarth/go-actor/printer"
)

func main() {
	p := printer.NewPrinterActor()
	p.Start()

	for i := 0; i < 100; i++ {
		message := printer.PrinterMessage{
			Tick:   time.Now(),
			Number: i,
		}
		p.Send(message)
	}

	select {}
}
