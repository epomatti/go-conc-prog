package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go receiveOrders(&wg) // "go" turns it async and managed by the scheduler
	wg.Wait()

	fmt.Println(orders)
}

func receiveOrders(wg *sync.WaitGroup) {
	for _, rawOrder := range rawOrder {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Println(err)
			continue
		}
		addOrder(newOrder)
	}
	wg.Done()
}

var rawOrder = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
