package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/OxygenTrader/go-life"
)

func main() {
	outDelay :=
		flag.Duration("out-delay", 100*time.Millisecond, "delay between frames")
	flag.Parse()

	text := "" +
		"!Name: Glider\n" +
		"!Author: Richard K. Guy\n" +
		"..........\n" +
		"..0.......\n" +
		"...0......\n" +
		".000......\n" +
		"..........\n" +
		"..........\n" +
		"..........\n" +
		"..........\n" +
		"..........\n" +
		"..........\n"
	field, err := life.ParseField(text)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Print(field.String())
		time.Sleep(*outDelay)

		field = field.NextField()
	}
}
