package main

import "strconv"

var (
	in = make(chan string)
)

func main() {
	worker(in)
}

func worker(in <-chan string) (chan int, chan error) {
	out := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		for msg := range in {
			i, err := strconv.Atoi(msg)
			if err != nil {
				errCh <- err
				return
			}
			out <- i
		}
	}()
	return out, errCh
}
