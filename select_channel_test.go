package belajar_golang_goroutine

import (
	"fmt"
	"testing"
)


func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	counter := 0

	go giveMeResponse(channel1)
	go giveMeResponse(channel2)

	for{
		select {
		case data := <- channel1:
			fmt.Println("Menerima data", data)
			counter++
		case data := <- channel2:
			fmt.Println("Menerima data", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}