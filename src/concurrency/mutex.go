package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const routines = 100
const delta = 1000

var m sync.Mutex

// sharedInt := 0 // implicit typing declaration is only allowed within functions
var sharedInt = 0
var realIncrementCount atomic.Int32

// has issues with lack of mut ex
// func incrementer() {
// 	for i := 0; i < delta; i++ {
// 		sharedInt++

// 		realIncrementCount.Add(1)
// 	}
// }

// uses mutexes for mut ex
// func incrementer() {
// 	for i := 0; i < delta; i++ {
// 		m.Lock()
// 		sharedInt++
// 		m.Unlock()

// 		realIncrementCount.Add(1)
// 	}
// }

func incrementer(num chan int) {
	for i := 0; i < delta; i++ {
		sharedIntCopy := <-num
		sharedIntCopy++
		num <- sharedIntCopy

		realIncrementCount.Add(1)
	}
}

func main() {
	realIncrementCount.Store(0)

	var wg sync.WaitGroup

	num := make(chan int, 1)

	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// incrementer()
			incrementer(num)
		}()
	}

	num <- sharedInt

	wg.Wait()

	sharedInt = <-num

	fmt.Println(sharedInt)
	fmt.Println(realIncrementCount.Load())
}
