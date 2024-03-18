package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(2500 * time.Millisecond)
	ch <- "Area 51"

}

func process2(ch2 chan string) {
	time.Sleep(2500 * time.Millisecond)
	ch2 <- "Area 52"
}

func main() {
	ch := make(chan string)
	ch2 := make(chan string)

	go process(ch)
	go process2(ch2)

	for {
		time.Sleep(1000 * time.Millisecond)

		// select çok kanaldan veri almak için ve sonsuz döngüdeyken kanaldan veri gelmez iken
		//default bir davranış sergilenebilmesi için de
		// buraya iki kanaldan da veri aynı anda gelirse receive sırasını rastgele seçer
		select {
		case v := <-ch:
			fmt.Println("received value from ch:", v)
			//return
		case v := <-ch2:
			fmt.Println("received value from ch2:", v)
			//return
		default:
			fmt.Println("no value received")
		}
	}
}
