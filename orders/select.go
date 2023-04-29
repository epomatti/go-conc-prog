package main

// go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
// loop:
// 	for {
// 		select {
// 		case order, ok := <-validOrderCh:
// 			if ok {
// 				fmt.Printf("Valid order received: %v\n", order)
// 			} else {
// 				break loop
// 			}
// 		case invalidOrder, ok := <-invalidOrderCh:
// 			if ok {
// 				fmt.Printf("Invalid order was received: %v. Issue: %v\n", invalidOrder.order, invalidOrder.err)
// 			} else {
// 				break loop
// 			}
// 		}
// 	}
// 	wg.Done()
// }(validOrderCh, invalidOrderCh)
