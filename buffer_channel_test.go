package belajar_golang_goroutine

import (
	"fmt"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	/*
	*jika tidak diisi jumlah buffer akan terjadi deadlock karena tidak ada yang menerima
	*defaultnya 1
	 */
	channel := make(chan string, 3)
	defer close(channel)

	/*
		Jika menggunakan buffer, ketika data dikirim ke buffer dan tidak diambil
		Maka data akan ada di buffer terus dan tidak terjadi deadlock
	*/
	channel <- "Arief"   //masuk ke buffer
	channel <- "Geming"  //masuk ke buffer
	channel <- "Ganteng" //masuk ke buffer
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	//fmt.Println(<-channel) //error deadlock karena buffer hanya 3
	fmt.Println("Selesai mengirim channel")
}
