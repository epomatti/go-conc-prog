package main

import "strconv"

var (
	in    = make(chan string)
	out   = make(chan int)
	errch = make(chan error, 1) // 💡 to avoid a potencial consumer not active
)

func main() {
	worker(in, out, errch)
}

func worker(in <-chan string, out chan<- int, errCh chan<- error) {
	for msg := range in {
		i, err := strconv.Atoi(msg)
		if err != nil {
			errCh <- err
			return
		}
		out <- i
	}
}
