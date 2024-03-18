package main

import "fmt"

// mutex için buradaki chan yapısını kullanabileceğimiz gibi sync içerisindeki mutex'i de kullanabiliriz
func main() {
	mutex := make(chan struct{}, 1)

	counter := 0

	incAtomic := func() {
		mutex <- struct{}{} // lock
		counter++
		<-mutex // unlock
	}

	increase10000 := func(done chan<- struct{}) {
		for i := 0; i < 10000; i++ {
			incAtomic()
		}

		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase10000(done)
	go increase10000(done)
	<-done
	<-done
	fmt.Println(counter)
}
