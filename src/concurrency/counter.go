package main

import (
	"fmt"
	"strconv"
)

func count(val *int, label string) {
	fmt.Printf("%v init val %v\n", label, *val)

	for i := 0; i < 1000000; i++ {
		*val = *val + 1
	}

	fmt.Printf("%v final val %v\n", label, *val)
}

func main() {
	val := 0
	for i := 0; i < 4; i++ {
		count(&val, "counter"+strconv.Itoa(i))
	}
	fmt.Printf("Done! val is %v\n", val)
}
