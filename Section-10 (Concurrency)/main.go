package main

import (
	"fmt"
	"time"
)

func greet(phase string) {
	fmt.Println(phase)
}

func slowGreet(phase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phase)
	doneChan <- true
}

func main() {

	// go greet("nice to meet you")
	// go greet("how are you")

	done := make(chan bool)

	go slowGreet("how...are....you", done)

	// go greet("i hope you are liking the course")

	<-done
}
