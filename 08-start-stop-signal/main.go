package main

import (
	"fmt"
	"time"
)

func main() {
	start := make(chan struct{})
	for i := 0; i < 10000; i++ {
		go func(k int) {
			<-start
			fmt.Println(k)
		}(i)
	}

	time.Sleep(3 * time.Second)

	// burada close sinyali yolladıktan sonra main goroutine sonlanacağı için
	// program yazabildiği kadar çıktıyı yazar hepsini yazamaz, main goroutine'inin
	// diğer goroutine'leri bitmesini beklesi için waitgroup kullanılması gerekir
	close(start)
}
