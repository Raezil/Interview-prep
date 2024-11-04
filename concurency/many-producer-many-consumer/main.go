package main

import (
	"fmt"
	"sync"
)

func Producer(wg *sync.WaitGroup, i int, p chan string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		p <- "Hello"
	}
}

func Consumer(wg *sync.WaitGroup, i int, p chan string) {
	defer wg.Done()

	for val := range p {
		text := fmt.Sprintf("Consumed %s by %d", val, i)
		fmt.Println(text)
	}
}

func main() {
	producerQ := make(chan string)
	wp := sync.WaitGroup{}
	wc := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wp.Add(1)
		go Producer(&wp, i, producerQ)
	}
	for i := 0; i < 2; i++ {
		wc.Add(1)
		go Consumer(&wc, i, producerQ)
	}
	wp.Wait()
	close(producerQ)
	wc.Wait()
}
