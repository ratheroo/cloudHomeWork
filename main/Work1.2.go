package main

import (
	"math/rand"
	"time"
)

func main() {
	c := make(chan int, 10)
	go produce(c)

	//consume
	func(cc <-chan int) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			n := <-cc
			println(n)
		}
	}(c)
}

func produce(c chan<- int) {
	timer := time.NewTicker(1 * time.Second)
	for _ = range timer.C {
		i := rand.Intn(9999)
		c <- i
	}
}
