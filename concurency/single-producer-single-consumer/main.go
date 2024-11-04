package main

import "fmt"

func Producer(p chan string) {
	p <- "Hello"
	close(p)
}

func Consumer(p chan string, done chan bool) {
	for val := range p {
		fmt.Println(val)
	}
	done <- true
}

func main() {
	p := make(chan string)
	done := make(chan bool)
	go Producer(p)
	go Consumer(p, done)
	<-done
}
