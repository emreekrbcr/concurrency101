package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
)

func main() {
	values := make([]byte, 32)
	//values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//boş struct bir alan kaplamadığı için genelde sinyal için kullanılır
	done := make(chan struct{})

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})

		fmt.Println("Sıralama işlemi tamamlandı.")
		done <- struct{}{}
	}()

	<-done

	fmt.Println(values[0], values[len(values)-1])
}
