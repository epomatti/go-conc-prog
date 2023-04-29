package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ch chan<- string) { // 💡 send-only channel
		ch <- "message"
	}(ch)
	go func(ch <-chan string) { // 💡 receive-only channel
		fmt.Println(<-ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
