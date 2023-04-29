package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for range time.Tick(500 * time.Millisecond) {
			fmt.Println("Tick!")
		}
		wg.Done()
	}()

	wg.Wait()
}
