package main

import "fmt"

func Producer(p chan string) {
	for i := 0; i < 10; i++ {
		p <- "Hello"
	}
	close(p)
}

func Consumer(i int, p chan string, done chan bool) {
	for val := range p {
		text := fmt.Sprintf("Consumed %s by %d", val, i)
		fmt.Println(text)
	}
	done <- true
}

func main() {
	p := make(chan string)
	done := make(chan bool)
	go Producer(p)
	for i := 0; i < 2; i++ {
		go Consumer(i, p, done)
	}
	<-done
}
