package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/OxygenTrader/go-life"
)

func main() {
	outDelay :=
		flag.Duration("out-delay", 100*time.Millisecond, "delay between frames")
	flag.Parse()

	textBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	field, err := life.ParseField(string(textBytes))
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Print(field.String())
		time.Sleep(*outDelay)

		field = field.NextField()
	}
}
