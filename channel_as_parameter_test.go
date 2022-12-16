package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func giveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Arief"
}

func TestChannelAsParameter(t *testing.T) {
	chennel := make(chan string)
	defer close(chennel)

	go giveMeResponse(chennel)

	data := <-chennel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
