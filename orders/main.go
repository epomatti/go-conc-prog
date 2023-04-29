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
	receivedOrdersCh := receiveOrders()
	validOrderCh, invalidOrderCh := validateOrders(receivedOrdersCh)

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

func validateOrders(in <-chan order) (<-chan order, <-chan invalidOrder) {
	out := make(chan order)
	errCh := make(chan invalidOrder, 1)
	go func() {
		for order := range in {
			if order.Quantity <= 0 {
				errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
			} else {
				out <- order
			}
		}
		close(out)
		close(errCh)
	}()
	return out, errCh
}

func receiveOrders() <-chan order {
	out := make(chan order)
	go func() {
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
	}()
	return out
}

var rawOrder = []string{
	`{"productCode": 1111, "quantity": -5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
