package main

import (
	"fmt"
	"time"
)

func main() {
	msgs := make(chan string) // HLchannel
	go func() {
		msgs <- "I need more coffee" // HLchannel
		close(msgs)
	}()

	go func() {
		msg := <-msgs // HLchannel
		fmt.Println(msg)
	}()
	time.Sleep(10 * time.Millisecond)
}
