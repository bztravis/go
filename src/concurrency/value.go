package main

import (
	"fmt"
	"strconv"
	"sync"
)

type value struct {
	adder   chan<- int       // send-only channel
	current <-chan int       // receive-only channel
	end     chan interface{} // empty interface channel
}

func makeValue() value {
	var result value

	adder := make(chan int)
	current := make(chan int)
	end := make(chan interface{})

	result.adder = adder
	result.current = current
	result.end = end

	go func() {
		val := 0

		for {
			select {
			case delta := <-adder:
				val += delta
			case current <- val:
			case <-end:
				return
			}
		}
	}()

	return result
}

func (v value) add(delta int) {
	v.adder <- delta
}

func (v value) get() int {
	return <-v.current
}

func (v value) done() {
	close(v.end)
}

func countWithVal(val value, label string) {
	fmt.Printf("%v init val %v\n", label, val.get())

	for i := 0; i < 1000000; i++ {
		val.add(1)
	}
	fmt.Printf("%v final val %v\n", label, val.get())

}

func main() {
	v := makeValue()
	defer v.done()

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) { // pass in the value to be captured as an argument, since it's not guaranteed to be the same value when the goroutine runs. This is fixed since Go v.1.22 https://tip.golang.org/doc/go1.22
			defer wg.Done()

			countWithVal(v, "counter"+strconv.Itoa(i))
		}(i)
	}

	wg.Wait()

	fmt.Printf("Done! val is %v\n", v.get())
}
