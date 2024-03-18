package main

import (
	"log"
	"math/rand"
	"time"
)

// counting semaphores

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Println("customer#", c, " enters the bar and waits")
	seat := <-bar // need a seat to drink
	log.Println("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Println("-- customer#", c, " frees seat#", seat)
	bar <- seat // free seat and leave the bar
}

func main() {
	rand.Seed(time.Now().UnixNano()) // programın her seferinde farklı rastgelelikte çalışmasını sağlar

	// the bar has 3 seats
	bar24x7 := make(Bar, 3)

	// place seats in the bar
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}
}
