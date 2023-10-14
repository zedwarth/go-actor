package main

import (
	"time"

	"github.com/zedwarth/go-actor/generator"
	"github.com/zedwarth/go-actor/printer"
)

func main() {
	p := printer.NewActor(10)
	p.Start()

	g := generator.NewActor(10)
	g.Start()

	ticker := time.Tick(time.Second)

	for tick := range ticker {
		g.Send(generator.Message{
			Printer: p,
			Tick:    tick,
		})
	}
}
