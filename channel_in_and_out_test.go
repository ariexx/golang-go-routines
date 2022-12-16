package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func OnlyIn(channel chan<- string) { //untuk mengirim data
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Arief" //memasukkan Muhammad Arief kedalam channel
}

func OnlyOut(channel <-chan string) { //untuk menerima
	data := <-channel //mengeluarkan channel ke dalam var data
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	chennel := make(chan string)
	defer close(chennel)

	go OnlyIn(chennel)
	go OnlyOut(chennel)

	time.Sleep(5 * time.Second)
}
