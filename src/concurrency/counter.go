package main

import (
	"fmt"
	"strconv"
	"sync"
)

func count(c chan int, label string) {
	val := <-c

	fmt.Printf("%v init val %v\n", label, val)

	for i := 0; i < 1000000; i++ {
		val++
	}

	fmt.Printf("%v final val %v\n", label, val)

	c <- val
}

func main() {
	val, num := 0, 4
	var wg sync.WaitGroup
	c := make(chan int, 1)

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) { // pass in the value to be captured as an argument, since it's not guaranteed to be the same value when the goroutine runs. This is fixed since Go v.1.22 https://tip.golang.org/doc/go1.22
			defer wg.Done()
			count(c, "counter"+strconv.Itoa(i))
		}(i)
	}

	c <- val
	wg.Wait()
	val = <-c

	fmt.Printf("Done! val is %v\n", val)
}
