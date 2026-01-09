package main

import (
	"fmt"
	"time"
)

func greet(phase string, doneChan chan bool) {
	fmt.Println(phase)
	doneChan <- true

}

func slowGreet(phase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(phase)
	doneChan <- true
	close(doneChan) // If we know the longest process that takes and we put close to that function it workds for The thrid way.

}

func main() {

	//This is the not scale way to do that one. For this reason we can use array of channels. This can be seen below part
	// done := make(chan bool)

	// go greet("nice to meet you", done)
	// go greet("how are you", done)

	// go slowGreet("how...are....you", done)

	// go greet("i hope you are liking the course", done)

	// <-done
	// <-done
	// <-done
	// <-done

	// ************************************************************************************************

	// dones := make([]chan bool, 4)

	// dones[0] = make(chan bool)
	// go greet("nice to meet you", dones[0])
	// dones[1] = make(chan bool)
	// go greet("how are you", dones[1])

	// dones[2] = make(chan bool)
	// go slowGreet("how...are....you", dones[2])
	// dones[3] = make(chan bool)
	// go greet("i hope you are liking the course", dones[3])

	// for _, done := range dones {
	// 	<-done
	// }

	// ************************************************************************************************

	done := make(chan bool)

	go greet("nice to meet you", done)
	go greet("how are you", done)

	go slowGreet("how...are....you", done)

	go greet("i hope you are liking the course", done)

	for range done {

	}

}
