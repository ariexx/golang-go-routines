package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)

	fmt.Println(time.Now())
	time2 := <-timer.C

	fmt.Println(time2)
}

func TestTimeAfter(t *testing.T) {
	channel := time.After(5 * time.Second)

	time2 := <-channel
	fmt.Println(time2)
}

func TestTimeAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Execute aftar 5 second")
		group.Done()
	})
	fmt.Println("Selesai") //dieksekusi terlebih dahulu

	group.Wait()
}
