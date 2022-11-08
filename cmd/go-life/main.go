package main

import (
	"fmt"
	"log"
	"time"

	"github.com/OxygenTrader/go-life"
)

func main() {
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
		time.Sleep(100 * time.Millisecond)

		field = field.NextField()
	}
}
