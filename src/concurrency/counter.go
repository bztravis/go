package main

import (
	"fmt"
	"strconv"
	"sync"
)

func count(val *int, label string) {
	fmt.Printf("%v init val %v\n", label, *val)

	for i := 0; i < 1000000; i++ {
		*val = *val + 1
	}

	fmt.Printf("%v final val %v\n", label, *val)
}

func main() {
	val, num := 0, 4
	var wg sync.WaitGroup

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) { // pass in the value to be captured as an argument, since it's not guaranteed to be the same value when the goroutine runs. This is fixed since Go v.1.22 https://tip.golang.org/doc/go1.22
			defer wg.Done()
			count(&val, "counter"+strconv.Itoa(i))
		}(i)
	}

	wg.Wait()

	fmt.Printf("Done! val is %v\n", val)
}
