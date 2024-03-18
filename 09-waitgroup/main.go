package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	start := make(chan struct{})

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func(k int) {
			defer wg.Done()
			<-start
			fmt.Println(k)
		}(i)
	}

	time.Sleep(3 * time.Second)

	close(start)

	wg.Wait()
}
