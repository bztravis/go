// Select statements are like switch statements but execute when one of the cases can proceed without being blocked.
// The interface only states that the cases are chosen randomly, not necessarily in order.

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		// case val, ok := <- quit:
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("waiting...")
		}
	}
}

// select statement runs when one of the cases can continue

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		// quit <- 0
		close(quit)
	}()
	fibonacci(c, quit)
}
