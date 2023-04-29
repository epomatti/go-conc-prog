package main

import "fmt"

func main() {
	ch := make(chan string, 1) // 💡 there is a memory price

	ch <- "message"

	fmt.Println(<-ch)
}
