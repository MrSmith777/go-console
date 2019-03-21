package main

import (
	"fmt"
	"github.com/DrSmithFr/go-console/pkg/output"
)

func main() {
	channel := make(chan string, 1)

	out := output.NewChanOutput(channel, true, nil)
	out.Writeln("Ceci est un test")

	msg := <-channel
	fmt.Print(msg)
}
