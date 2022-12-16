package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Arief geming" //terblock karna tidak ada yang mengambil, script dibawah akan terblock, waiting terus
		fmt.Println("Selesai mengirim data ke channel")
	}()
	fmt.Println(<-channel)
	time.Sleep(5 * time.Second)
}
