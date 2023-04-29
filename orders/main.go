package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receivedOrdersCh = make(chan order)
	var validOrderCh = make(chan order)
	var invalidOrderCh = make(chan invalidOrder)
	go receiveOrders(receivedOrdersCh) // 💡 Adding "go" makes it async and managed by the scheduler
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)

	wg.Add(1) // 💡 Channels are blocking, so Add(1) is enough

	go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Printf("Valid order received: %v\n", order)
				} else {
					break loop
				}
			case invalidOrder, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order was received: %v. Issue: %v\n", invalidOrder.order, invalidOrder.err)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}

func validateOrders(in <-chan order, out chan<- order, errCh chan<- invalidOrder) {
	for order := range in {
		if order.Quantity <= 0 {
			errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
		} else {
			out <- order
		}
	}
	close(out)
	close(errCh)
}

func receiveOrders(out chan<- order) {
	for _, rawOrder := range rawOrder {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		out <- newOrder
	}
	close(out)
}

var rawOrder = []string{
	`{"productCode": 1111, "quantity": -5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
