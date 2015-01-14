package main

import (
	"fmt"
	"time"
)

func main() {

	go func() { // HLgoroutine
		for i := 0; i < 5; i++ {
			fmt.Println("Welcome!")
		}
	}()

	time.Sleep(10 * time.Millisecond)
}
