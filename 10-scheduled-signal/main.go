package main

import (
	"fmt"
	"time"
)

func AfterDuration(d time.Duration) <-chan struct{} {
	ch := make(chan struct{}, 1)

	go func() {
		time.Sleep(d)
		ch <- struct{}{}
	}()

	return ch
}

func main() {
	fmt.Println("Hi")
	<-AfterDuration(time.Second)
	fmt.Println("Hello")
	<-AfterDuration(time.Second)
	fmt.Println("Bye")

	<-time.After(time.Second) // Kendi yazdığımız fonksiyonun hazırı
	fmt.Println("Last")
}
