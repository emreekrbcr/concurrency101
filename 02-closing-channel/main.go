package main

import "fmt"

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		if i == 4 {
			close(ch)
			break
		}

		ch <- i
	}

	//close(ch)
}

func main() {
	ch := make(chan int)

	go producer(ch)

	for {
		v, ok := <-ch

		if ok == false {
			fmt.Println("kırdı")
			break
		}

		fmt.Println("Received ", v, ok)
	}
}
